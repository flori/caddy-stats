package stats

import (
	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

func setup(c *caddy.Controller) error {
	config := httpserver.GetConfig(c)
	args := c.RemainingArgs()
	ru := ""
	if len(args) > 1 {
		ru = args[1]
	}
	middleware := func(next httpserver.Handler) httpserver.Handler {
		return StatsHandler{RedisURL: ru, Next: next}
	}
	config.AddMiddleware(middleware)
	return nil
}

func init() {
	caddy.RegisterPlugin("stats", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
	httpserver.RegisterDevDirective("stats", "log")
}
