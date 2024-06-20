package sight_middleware

import (
	"github.com/labstack/echo/v4"
	"regexp"
)

var validPath = regexp.MustCompile(`^/v[0-9]+/.*`)
var appName, env string

// ApiMetrics logs API request
func ApiMetrics(app, e string) echo.MiddlewareFunc {
	appName = app
	env = e
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			routeName := c.Path()
			//start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			if validPath.MatchString(routeName) {
				//duration := time.Now().Sub(start)
				//res := c.Response()
				//if err := lolo_aws.PutApiMetric(appName+"-"+env, routeName, float64(duration.Milliseconds()), map[string]string{
				//	"app":    appName,
				//	"route":  routeName,
				//	"status": strconv.Itoa(res.Status),
				//}); err != nil {
				//	log.Error().Err(err).Msg("ApiMetrics Client err")
				//}
			}
			return
		}
	}
}
