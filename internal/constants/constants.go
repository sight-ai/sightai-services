package constants

import "github.com/shopspring/decimal"

const (
	OKResponse        = "OK"
	AdminAddress      = "0x740871b43ff559cc6d31b7ad9744cbf2a22bb9ae"
	SightTokenDecimal = 18
	DefaultGatewayID  = 1
)

var (
	NewUserBonus, _ = decimal.NewFromString("500")
)
