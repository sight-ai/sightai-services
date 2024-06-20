package sight_middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	// https://github.com/grpc-ecosystem/grpc-gateway#mapping-grpc-to-http
	// HTTP headers with this prefix are set as GRPC metadata on handler side.
	grpcgwHeaderPrefix     = "Grpc-Metadata-"
	grpcgwHeaderXRequestID = grpcgwHeaderPrefix + echo.HeaderXRequestID
)

var (
	requestIDConfig = func() middleware.RequestIDConfig {
		cfg := middleware.DefaultRequestIDConfig
		cfg.Generator = func() string {
			str, _ := uuid.NewRandom()
			return str.String()
		}
		return cfg
	}()
)

// AddRequestID add a random UUID to the request
func AddRequestID() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if requestIDConfig.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()

			rid := req.Header.Get(echo.HeaderXRequestID)
			if rid == "" {
				rid = requestIDConfig.Generator()
				res.Header().Set(echo.HeaderXRequestID, rid)
				req.Header.Set(echo.HeaderXRequestID, rid)
				c.Set(echo.HeaderXRequestID, rid)
			} else {
				c.Set(echo.HeaderXRequestID, rid)
			}

			return next(c)
		}
	}
}

func AddRequestIDHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(grpcgwHeaderXRequestID) == "" {
			r.Header.Set(grpcgwHeaderXRequestID, requestIDConfig.Generator())
		}
		next.ServeHTTP(w, r)
	})
}
