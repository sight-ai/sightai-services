package aes_cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"strings"

	pkgerrors "github.com/pkg/errors"
)

type AESCipher struct {
	block cipher.Block
}

func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}

	return value
}

func removeBase64Padding(value string) string {
	return strings.Replace(value, "=", "", -1)
}

func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func NewAESCipher(key []byte) (*AESCipher, error) {
	cipher := &AESCipher{}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cipher.block = block
	return cipher, nil
}

func (c *AESCipher) Encrypt(text string) (string, error) {
	msg := pad([]byte(text))
	cipherText := make([]byte, aes.BlockSize+len(msg))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(c.block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(msg))
	finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(cipherText))
	return finalMsg, nil
}

func (c *AESCipher) Decrypt(text string) (string, error) {
	decodedMsg, err := base64.URLEncoding.DecodeString(addBase64Padding(text))
	if err != nil {
		return "", err
	}

	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("blocksize must be multipe of decoded message length")
	}

	iv := decodedMsg[:aes.BlockSize]
	msg := decodedMsg[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(c.block, iv)
	cfb.XORKeyStream(msg, msg)

	unpadMsg, err := unpad(msg)
	if err != nil {
		return "", err
	}

	return string(unpadMsg), nil
}

func (c *AESCipher) EncryptStruct(v interface{}) (string, error) {
	bz, err := json.Marshal(v)
	if err != nil {
		return "", pkgerrors.Wrap(err, "failed to marshal arg to json")
	}
	return c.Encrypt(string(bz))
}

func (c *AESCipher) DecryptToStruct(enc string, output interface{}) error {
	s, err := c.Decrypt(enc)
	if err != nil {
		return pkgerrors.Wrap(err, "failed to decrypt arg")
	}

	if err := json.Unmarshal([]byte(s), output); err != nil {
		return pkgerrors.Wrap(err, "failed to unmarshal decrypted value to given struct")
	}
	return err
}
