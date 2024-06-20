package itest

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Deposit(t *testing.T) {
	resp, err := Deposit(JwtTokenUID1, "0xb6D402c622E130738dd40872D51c70515381C33c", "10000")
	require.Nil(t, err)
	require.NotNil(t, resp)
}
