package rest

import (
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_PageToken(t *testing.T) {
	page := rand.Int63()
	token, err := EncryptPageToken(&PageToken{Page: page})
	assert.Nil(t, err)
	log.Info().Msgf(token)
	res, err := DecryptPageToken(token)
	assert.Nil(t, err)
	assert.Equal(t, page, res.Page)
}
