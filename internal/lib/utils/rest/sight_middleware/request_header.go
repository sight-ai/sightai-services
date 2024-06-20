package sight_middleware

import (
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/labstack/echo/v4"
)

// RequestHeader add params header
func RequestHeader(args ...string) echo.MiddlewareFunc {
	if len(args) == 0 || len(args)%2 != 0 {
		log.Error().Msgf("failed to add wrapHeader middleware, # of args is %v", len(args))
		return nil
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			for i := 0; i < len(args); i += 2 {
				req.Header.Set(args[i], args[i+1])
				c.Set(args[i], args[i+1])
			}
			return next(c)
		}
	}
}
