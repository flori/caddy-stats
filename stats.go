package stats

import (
	"log"
	"net/http"
	"os"

	"github.com/flori/caddy-stats/log_json"
	"github.com/go-redis/redis"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

const (
	REDIS_URL_VARIABLE    = "REDIS_URL"
	LOG_MAX_SIZE_VARIABLE = "LOG_MAX_SIZE"
	LOG_SET               = "LOG"
)

type StatsHandler struct {
	RedisURL   string
	LogMaxSize int
	Next       httpserver.Handler
}

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
		log.Panic(err)
	}
	return redis.NewClient(options)
}

func addToLog(c *redis.Client, l log_json.LogJSON, logMaxSize int) {
	zadd := c.ZAdd(LOG_SET, redis.Z{float64(l.Time.Unix()), l.String()})
	if zadd.Err() != nil {
		log.Panic(zadd.Err())
	}
	zrr := c.ZRemRangeByRank(LOG_SET, 0, int64(-logMaxSize-1))
	if zrr.Err() != nil {
		log.Panic(zrr.Err())
	}
}

func (h StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	c := newRedisClient(h)
	defer c.Close()
	l := log_json.NewLogJSON(r)
	incr := c.Incr("HIT:" + l.Target())
	if incr.Err() != nil {
		log.Panic(incr.Err())
	}
	addToLog(c, l, h.LogMaxSize)
	return h.Next.ServeHTTP(w, r)
}
