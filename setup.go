package stats

import (
	"log"
	"os"
	"strconv"

	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

var lms = 10000

func logMaxSize() int {
	size, ok := os.LookupEnv(LOG_MAX_SIZE_VARIABLE)
	if !ok {
		goto Default
	}
	if s, err := strconv.Atoi(size); err == nil {
		if s >= 0 {
			return s
		}
	}
Default:
	return 10000
}

func setup(c *caddy.Controller) error {
	config := httpserver.GetConfig(c)
	args := c.RemainingArgs()
	ru := ""
	if len(args) > 1 {
		ru = args[1]
	}
	middleware := func(next httpserver.Handler) httpserver.Handler {
		return StatsHandler{RedisURL: ru, LogMaxSize: lms, Next: next}
	}
	config.AddMiddleware(middleware)
	return nil
}

func init() {
	lms = logMaxSize()
	log.Printf("Setting up caddy-stats middleware with LOG_MAX_SIZE = %d", lms)
	caddy.RegisterPlugin("stats", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
	httpserver.RegisterDevDirective("stats", "log")
}
