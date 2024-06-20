package rest

import (
	"errors"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/aes_cipher"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/foolin/mixer"
	"github.com/labstack/echo/v4"
	"reflect"
	"strings"
)

type DTO interface{}

var aesCipher *aes_cipher.AESCipher
var mixerPassword string

type PageToken struct {
	Page int64 `json:"page"`
}

func initialize(key, password string) {
	var err error
	aesCipher, err = aes_cipher.NewAESCipher([]byte(key))
	if err != nil {
		log.Fatal().Msgf("init rest module error: aesCipher failed")
	}
	mixerPassword = password
}

func parseRequest(c echo.Context, i DTO) error {
	if c == nil {
		log.Error(c).Msg("empty context")
		return errors.New("context is empty")
	}
	err := c.Bind(i)
	if err != nil {
		log.Error(c).Msgf("failed to parse request, err=\"%v\"", err)
		return errors.New("failed to parse request")
	}
	if err = Validate(i); err != nil {
		log.Error(c).Msgf("invalid request=%#v, err=\"%v\"", i, strings.ReplaceAll(err.Error(), "\n", ", "))
		return errors.New(strings.ReplaceAll(err.Error(), "\n", ", "))
	}
	return nil
}

// GetRequest is to obtain the http request from echo context
func GetRequest(c echo.Context, i DTO) (interface{}, error) {
	dtoVal := reflect.New(reflect.TypeOf(i).Elem())
	dto, _ := dtoVal.Interface().(DTO)

	err := parseRequest(c, dto)
	if err != nil {
		return nil, err
	}

	return dto.(DTO), err
}

// GenRequest is to obtain the http request from echo context
func GenRequest(c echo.Context, i DTO) error {
	dto, _ := i.(DTO)
	err := parseRequest(c, dto)
	if err != nil {
		return err
	}

	return nil
}

// DecryptPageToken using aes encryption
func DecryptPageToken(token string) (*PageToken, error) {
	var res PageToken
	err := aesCipher.DecryptToStruct(token, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// EncryptPageToken using aes encryption
func EncryptPageToken(pageToken *PageToken) (string, error) {
	return aesCipher.EncryptStruct(pageToken)
}

// DecryptID using aes encryption
func DecryptID(token string) (uint, error) {
	id, err := mixer.DecodeID(mixerPassword, token)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

// EncryptID using aes encryption
func EncryptID(id uint) string {
	return mixer.EncodeIDPadding(mixerPassword, uint64(id), 10)
}
