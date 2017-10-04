package stats

import (
	"github.com/flori/caddy-stats/log_json"
	"github.com/go-redis/redis"
	"github.com/mholt/caddy/caddyhttp/httpserver"
	"log"
	"net/http"
	"os"
)

type StatsHandler struct {
	RedisURL string
	Next     httpserver.Handler
}

const REDIS_URL_VARIABLE = "REDIS_URL"

func redisURL() string {
	url, ok := os.LookupEnv(REDIS_URL_VARIABLE)
	if !ok {
		url = "redis://localhost:6379"
	}
	return url
}

func newRedisClient(h StatsHandler) *redis.Client {
	ru := h.RedisURL
	if ru == "" {
		ru = redisURL()
	}
	options, err := redis.ParseURL(ru)
	if err != nil {
		log.Fatal(err)
	}
	return redis.NewClient(options)
}

func (h StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	c := newRedisClient(h)
	l := log_json.NewLogJSON(r)
	incr := c.Incr("HIT:" + l.Target())
	if incr.Err() != nil {
		log.Fatal(incr.Err())
	}
	zadd := c.ZAdd("LOG", redis.Z{float64(l.Time.Unix()), l.String()})
	if zadd.Err() != nil {
		log.Fatal(zadd.Err())
	}
	return h.Next.ServeHTTP(w, r)
}
