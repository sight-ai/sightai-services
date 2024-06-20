package sight_middleware

import (
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	// MaxResponseSize - 100kb
	MaxResponseSize = 100 * 1024
)

// ResponseDump of API request
func ResponseDump(env string) echo.MiddlewareFunc {
	h := func(c echo.Context, reqBody, resBody []byte) {
		if len(resBody) > 0 && len(resBody) < MaxResponseSize {
			log.Info(c).
				Str("response", string(compactJson(resBody))).
				Msg("response_dump")
		}
	}
	if env == "prod" {
		h = func(c echo.Context, reqBody, resBody []byte) {
			if len(resBody) > 0 && len(resBody) < MaxResponseSize {
				log.Info(c).
					RawJSON("response", compactJson(resBody)).
					Msg("response_dump")
			}
		}
	}
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: func(c echo.Context) bool {
			if c.Request().URL.Path == "/health-check" {
				return true
			}
			return false
		},
		Handler: h,
	})
}
