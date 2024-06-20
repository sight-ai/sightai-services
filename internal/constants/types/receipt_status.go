package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type ReceiptStatus string

const (
	ReceiptStatusFinished = ReceiptStatus("finished")
	ReceiptStatusPaid     = ReceiptStatus("paid")
)

func (t ReceiptStatus) IsValid() error {
	switch t {
	case ReceiptStatusFinished, ReceiptStatusPaid:
		return nil
	}
	return errors.New("invalid OrderBy")
}

func (t *ReceiptStatus) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	vBytes, ok := value.([]byte)
	if !ok {
		return errors.New("except byte slice")
	}

	*t = ReceiptStatus(vBytes)
	if err := t.IsValid(); err != nil {
		return err
	}
	return nil
}

func (t ReceiptStatus) Value() (driver.Value, error) {
	if t == "" {
		return nil, nil
	}
	if err := t.IsValid(); err != nil {
		return nil, err
	}
	return []byte(t), nil
}

func NewReceiptStatusFromString(s string) (*ReceiptStatus, error) {
	status := ReceiptStatus(s)
	switch status {
	case ReceiptStatusFinished, ReceiptStatusPaid:
	default:
		return nil, fmt.Errorf("invalid ReceiptStatus: %v", s)
	}
	return &status, nil
}

func (t ReceiptStatus) String() string {
	return string(t)
}
