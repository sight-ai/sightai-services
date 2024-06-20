package jwt_auth

import (
	"crypto/rsa"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	JwtMapUIDField  = "uid"
	JwtMapAddrField = "address"
	JwtMapExpField  = "exp"
)

type JWT struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var JwtKeys *JWT

func LazyLoadJwt() {
	if JwtKeys == nil {
		prvKey, _ := os.ReadFile(config.Cfg.JwtPrv)

		pubKey, _ := os.ReadFile(config.Cfg.JwtPub)

		privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(prvKey)
		if err != nil {
			log.Fatalln(err)
		}

		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
		if err != nil {
			log.Fatalln(err)
		}

		JwtKeys = &JWT{
			privateKey,
			publicKey,
		}
	}
}

func GenerateJwtFromAccount(account *entities.Account) (string, error) {
	LazyLoadJwt()

	// token with claims
	claims := jwt.MapClaims{}
	claims[JwtMapUIDField] = account.ID
	claims[JwtMapAddrField] = account.Address
	claims[JwtMapExpField] = time.Now().Add(time.Hour * 24 * 365).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	t, err := token.SignedString(JwtKeys.PrivateKey)
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetAccountID(jwtToken string) (uint, error) {
	LazyLoadJwt()

	if strings.HasPrefix(jwtToken, "Bearer") {
		jwtToken = strings.TrimSpace(jwtToken[len("Bearer"):])
	}
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(
		jwtToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return JwtKeys.PublicKey, nil
		},
	)

	if err != nil {
		return 0, err
	}
	return uint(claims[JwtMapUIDField].(float64)), nil
}
