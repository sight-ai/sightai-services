package aes_cipher

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestS struct {
	X string
	Y int
}

func TestEncryptDecryptStruct(t *testing.T) {
	a, err := NewAESCipher([]byte("68566D5971337436"))
	require.NoError(t, err)

	// encrypt integer string
	enc, err := a.Encrypt("1")
	require.NoError(t, err)
	require.NotEmpty(t, enc)

	fmt.Println(enc)

	// decrypt
	intStrOut, err := a.Decrypt(enc)
	require.NoError(t, err)

	require.Equal(t, "1", intStrOut)

	// encrypt struct
	enc, err = a.EncryptStruct(&TestS{"test", 1})
	require.NoError(t, err)
	require.NotEmpty(t, enc)

	// decrypt
	structOut := &TestS{}
	require.NoError(t, a.DecryptToStruct(enc, structOut))

	require.Equal(t, "test", structOut.X)
	require.True(t, 1 == structOut.Y)
}
