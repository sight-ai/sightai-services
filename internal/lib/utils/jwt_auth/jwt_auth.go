package jwt_auth

import (
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

const (
	JwtMapUIDField  = "uid"
	JwtMapAddrField = "address"
	JwtMapExpField  = "exp"
)

//type JWT struct {
//	privateKey *rsa.PrivateKey
//	publicKey  *rsa.PublicKey
//}
//
//var jwtKeys *JWT
//
//func lazyLoadJwt() {
//	if jwtKeys == nil {
//		prvKey, err := os.ReadFile(config.Cfg.JwtPrv)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		pubKey, err := os.ReadFile(config.Cfg.JwtPub)
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(prvKey)
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		jwtKeys = &JWT{
//			privateKey,
//			publicKey,
//		}
//	}
//}

func GenerateJwtFromAccount(account *entities.Account) (string, error) {
	//lazyLoadJwt()

	// token with claims
	claims := jwt.MapClaims{}
	claims[JwtMapUIDField] = account.ID
	claims[JwtMapAddrField] = account.Address
	claims[JwtMapExpField] = time.Now().Add(time.Hour * 24 * 365).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	t, err := token.SignedString([]byte(config.Cfg.JwtSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetAccountID(jwtToken string) (uint, error) {
	//lazyLoadJwt()

	if strings.HasPrefix(jwtToken, "Bearer") {
		jwtToken = strings.TrimSpace(jwtToken[len("Bearer"):])
	}
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(
		jwtToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Cfg.JwtSecret), nil
		},
	)
	if err != nil {
		return 0, err
	}
	return uint(claims[JwtMapUIDField].(float64)), nil
}
