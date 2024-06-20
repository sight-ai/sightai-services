package sight_middleware

import (
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/jwt_auth"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

const (
	JwtMapUIDField = "uid"
)

func Jwt() echo.MiddlewareFunc {
	jwt_auth.LazyLoadJwt()
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    jwt_auth.JwtKeys.PublicKey,
		TokenLookup:   fmt.Sprintf("header:%s", rest.XUserToken),
		SigningMethod: jwt.SigningMethodRS256.Name,
	})
}

func GetUserIDFromJwt(jwtToken string) (uint, error) {
	jwt_auth.LazyLoadJwt()
	if strings.HasPrefix(jwtToken, "Bearer") {
		jwtToken = strings.TrimSpace(jwtToken[len("Bearer"):])
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		jwtToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwt_auth.JwtKeys.PublicKey, nil
		},
	)
	if err != nil {
		return 0, err
	}
	return uint(claims[JwtMapUIDField].(float64)), nil
}
