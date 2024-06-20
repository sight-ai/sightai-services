package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type AccountRole string

const (
	AccountRoleUser    = AccountRole("user")
	AccountRoleGateway = AccountRole("gateway")
	AccountRoleAdmin   = AccountRole("admin")
)

func (s AccountRole) IsValid() error {
	switch s {
	case AccountRoleUser, AccountRoleGateway, AccountRoleAdmin:
		return nil
	}
	return errors.New("invalid AccountRole")
}

func (s *AccountRole) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	vBytes, ok := value.([]byte)
	if !ok {
		return errors.New("except byte slice")
	}

	*s = AccountRole(vBytes)
	if err := s.IsValid(); err != nil {
		return err
	}
	return nil
}

func (s AccountRole) Value() (driver.Value, error) {
	if s == "" {
		return nil, nil
	}
	if err := s.IsValid(); err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func NewLiquidityTypeFromString(s string) (*AccountRole, error) {
	liquidityType := AccountRole(s)
	switch liquidityType {
	case AccountRoleUser, AccountRoleGateway, AccountRoleAdmin:
	default:
		return nil, fmt.Errorf("invalid AccountRole: %v", s)
	}
	return &liquidityType, nil
}

func (s AccountRole) String() string {
	return string(s)
}
