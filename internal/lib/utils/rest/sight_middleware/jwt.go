package sight_middleware

import (
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

const (
	JwtMapUIDField = "uid"
)

var jwtSecret string

func Jwt(secret string) echo.MiddlewareFunc {
	jwtSecret = secret
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(secret),
		TokenLookup:   fmt.Sprintf("header:%s", rest.XUserToken),
		SigningMethod: jwt.SigningMethodHS256.Name,
	})
}

func GetUserIDFromJwt(jwtToken string) (uint, error) {
	if strings.HasPrefix(jwtToken, "Bearer") {
		jwtToken = strings.TrimSpace(jwtToken[len("Bearer"):])
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		jwtToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		},
	)
	if err != nil {
		return 0, err
	}
	return uint(claims[JwtMapUIDField].(float64)), nil
}
