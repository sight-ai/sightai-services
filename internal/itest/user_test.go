package itest

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_RegisterDevice(t *testing.T) {
	resp, err := SignIn("0xA")
	require.Nil(t, err)
	require.NotNil(t, resp)
}
