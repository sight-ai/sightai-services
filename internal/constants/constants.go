package constants

import "github.com/shopspring/decimal"

const (
	OKResponse        = "OK"
	SightTokenDecimal = 18
	DefaultGatewayID  = 1
)

var (
	NewUserBonus, _ = decimal.NewFromString("500")
)
