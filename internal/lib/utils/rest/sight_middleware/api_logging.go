package sight_middleware

import (
	"bytes"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/color"
	"github.com/valyala/fasttemplate"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ApiLoggingJson logs API request
func ApiLoggingJson() echo.MiddlewareFunc {
	cfg := middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			if c.Request().URL.Path == "/health-check" {
				return true
			}
			return false
		},
		Format: `{"time":"${time_rfc3339_nano}","req_id":"${req_id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}",` +
			`"bytes_in":${bytes_in},"bytes_out":${bytes_out},"client_key":"${client_key}",` +
			`"user_id":"${user_id}"}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}
	return loggerWithConfig(cfg)
}

// ApiLogging logs API request
func ApiLogging(args ...string) echo.MiddlewareFunc {
	cfg := middleware.DefaultLoggerConfig

	cfg.Skipper = func(c echo.Context) bool {
		if c.Request().URL.Path == "/health-check" {
			return true
		}
		return false
	}
	if len(args) > 0 {
		cfg.Format = `${time_rfc3339} ` + strings.ToLower(args[0]) + ` ${status} ${method} ${uri} latency=${latency_human}, ` +
			`bytesIn=${bytes_in}, bytesOut=${bytes_out}, req_id=${req_id}, from=${remote_ip}, fromService=${from_service}, err=${error}, ua=${user_agent}, ` +
			`client_key=${client_key}, user_id=${user_id}` + "\n"
	} else {
		cfg.Format = `${time_rfc3339} API ${status} ${method} ${uri} latency=${latency_human}, ` +
			`bytesIn=${bytes_in}, bytesOut=${bytes_out}, req_id=${req_id}, from=${remote_ip}, fromService=${from_service} err=${error}, ua=${user_agent}, ` +
			`client_key=${client_key}, user_id=${user_id}` + "\n"
	}
	return loggerWithConfig(cfg, args...)
}

// LoggerWithConfig returns a Logger middleware with config.
func loggerWithConfig(config middleware.LoggerConfig, args ...string) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultLoggerConfig.Skipper
	}
	if config.Format == "" {
		config.Format = middleware.DefaultLoggerConfig.Format
	}
	if config.Output == nil {
		config.Output = middleware.DefaultLoggerConfig.Output
	}

	template := fasttemplate.New(config.Format, "${", "}")
	colorer := color.New()
	colorer.SetOutput(config.Output)
	pool := &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 256))
		},
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			buf := pool.Get().(*bytes.Buffer)
			buf.Reset()
			defer pool.Put(buf)

			if _, err = template.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
				switch tag {
				case "time_unix":
					return buf.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
				case "time_unix_nano":
					return buf.WriteString(strconv.FormatInt(time.Now().UnixNano(), 10))
				case "time_rfc3339":
					return buf.WriteString(colorer.Grey(time.Now().Format(time.RFC3339)))
				case "time_rfc3339_nano":
					return buf.WriteString(time.Now().Format(time.RFC3339Nano))
				case "time_custom":
					return buf.WriteString(time.Now().Format(config.CustomTimeFormat))
				case "req_id":
					id := req.Header.Get(echo.HeaderXRequestID)
					if id == "" {
						id = res.Header().Get(echo.HeaderXRequestID)
					}
					return buf.WriteString(id)
				case "remote_ip":
					return buf.WriteString(c.RealIP())
				case "host":
					return buf.WriteString(req.Host)
				case "uri":
					return buf.WriteString(req.RequestURI)
				case "method":
					return buf.WriteString(req.Method)
				case "path":
					p := req.URL.Path
					if p == "" {
						p = "/"
					}
					return buf.WriteString(p)
				case "protocol":
					return buf.WriteString(req.Proto)
				case "referer":
					return buf.WriteString(req.Referer())
				case "user_agent":
					return buf.WriteString(req.UserAgent())
				case "status":
					n := res.Status
					s := colorer.Green(n)
					switch {
					case n >= 500:
						s = colorer.RedBg(n)
					case n >= 400:
						s = colorer.YellowBg(n)
					case n >= 300:
						s = colorer.Cyan(n)
					}
					return buf.WriteString(s)
				case "error":
					if err != nil {
						return buf.WriteString(err.Error())
					}
				case "latency":
					l := stop.Sub(start)
					return buf.WriteString(strconv.FormatInt(int64(l), 10))
				case "latency_human":
					return buf.WriteString(stop.Sub(start).String())
				case "bytes_in":
					cl := req.Header.Get(echo.HeaderContentLength)
					if cl == "" {
						cl = "0"
					}
					return buf.WriteString(cl)
				case "bytes_out":
					return buf.WriteString(strconv.FormatInt(res.Size, 10))
				case "user_id":
					uid, _ := GetUserIDFromJwt(req.Header.Get(rest.XUserToken))
					return buf.WriteString(strconv.FormatInt(int64(uid), 10))
				default:
					switch {
					case strings.HasPrefix(tag, "header:"):
						return buf.Write([]byte(c.Request().Header.Get(tag[7:])))
					case strings.HasPrefix(tag, "query:"):
						return buf.Write([]byte(c.QueryParam(tag[6:])))
					case strings.HasPrefix(tag, "form:"):
						return buf.Write([]byte(c.FormValue(tag[5:])))
					case strings.HasPrefix(tag, "cookie:"):
						cookie, err := c.Cookie(tag[7:])
						if err == nil {
							return buf.Write([]byte(cookie.Value))
						}
					}
				}
				return 0, nil
			}); err != nil {
				return
			}

			if config.Output == nil {
				_, err = c.Logger().Output().Write(buf.Bytes())
				return
			}
			_, err = config.Output.Write(buf.Bytes())
			return
		}
	}
}
