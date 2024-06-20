package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type TransactionType string

const (
	TransactionTypeDeposit       = TransactionType("deposit")
	TransactionTypeDepositEvent  = TransactionType("depositevent")
	TransactionTypeWithdraw      = TransactionType("withdraw")
	TransactionTypeWithdrawEvent = TransactionType("withdrawevent")
	TransactionTypeAllowance     = TransactionType("allowance")
	TransactionTypePay           = TransactionType("pay")
	TransactionTypeReceive       = TransactionType("receive")
)

func (s TransactionType) IsValid() error {
	switch s {
	case TransactionTypeDeposit, TransactionTypeDepositEvent, TransactionTypeWithdraw, TransactionTypeWithdrawEvent, TransactionTypeAllowance, TransactionTypePay, TransactionTypeReceive:
		return nil
	}
	return errors.New("invalid TransactionType")
}

func (s *TransactionType) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	vBytes, ok := value.([]byte)
	if !ok {
		return errors.New("except byte slice")
	}

	*s = TransactionType(vBytes)
	if err := s.IsValid(); err != nil {
		return err
	}
	return nil
}

func (s TransactionType) Value() (driver.Value, error) {
	if s == "" {
		return nil, nil
	}
	if err := s.IsValid(); err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func NewTransactionTypeFromString(s string) (*TransactionType, error) {
	billType := TransactionType(s)
	switch billType {
	case TransactionTypeDeposit, TransactionTypeDepositEvent, TransactionTypeWithdraw, TransactionTypeWithdrawEvent, TransactionTypeAllowance, TransactionTypePay, TransactionTypeReceive:
	default:
		return nil, fmt.Errorf("invalid TransactionType: %v", s)
	}
	return &billType, nil
}

func (s TransactionType) String() string {
	return string(s)
}
