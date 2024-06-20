package sight_middleware

import (
	"bytes"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

// RequestDump of API request
func RequestDump(env string) echo.MiddlewareFunc {
	h := func(c echo.Context, reqBody []byte) {
		if len(reqBody) > 0 {
			uid, _ := GetUserIDFromJwt(c.Request().Header.Get(rest.XUserToken))
			log.Info(c).
				Str("host", c.Request().Host).
				Str("method", c.Request().Method).
				Str("uri", c.Request().RequestURI).
				Str("request", string(compactJson(reqBody))).
				Uint("user_id", uid).
				Msg("request_dump")
		}
	}
	if env == "prod" {
		h = func(c echo.Context, reqBody []byte) {
			uid, _ := GetUserIDFromJwt(c.Request().Header.Get(rest.XUserToken))
			if len(reqBody) > 0 {
				log.Info(c).
					Str("host", c.Request().Host).
					Str("method", c.Request().Method).
					Str("uri", c.Request().RequestURI).
					RawJSON("request", compactJson(reqBody)).
					Uint("user_id", uid).
					Msg("request_dump")
			}
		}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			// Request
			reqBody := []byte{}
			if c.Request().Body != nil { // Read
				reqBody, _ = ioutil.ReadAll(c.Request().Body)
			}
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset

			h(c, reqBody)

			if err = next(c); err != nil {
				c.Error(err)
			}

			return
		}
	}
}
