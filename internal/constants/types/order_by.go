package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type OrderBy string

const (
	OrderByCreateTimeNew = OrderBy("new")
	OrderByCreateTimeOld = OrderBy("old")
)

func (t OrderBy) IsValid() error {
	switch t {
	case OrderByCreateTimeNew, OrderByCreateTimeOld:
		return nil
	}
	return errors.New("invalid OrderBy")
}

func (t *OrderBy) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	vBytes, ok := value.([]byte)
	if !ok {
		return errors.New("except byte slice")
	}

	*t = OrderBy(vBytes)
	if err := t.IsValid(); err != nil {
		return err
	}
	return nil
}

func (t OrderBy) Value() (driver.Value, error) {
	if t == "" {
		return nil, nil
	}
	if err := t.IsValid(); err != nil {
		return nil, err
	}
	return []byte(t), nil
}

func NewOrderByFromString(s string) (*OrderBy, error) {
	status := OrderBy(s)
	switch status {
	case OrderByCreateTimeNew, OrderByCreateTimeOld:
	default:
		return nil, fmt.Errorf("invalid OrderBy: %v", s)
	}
	return &status, nil
}

func (t OrderBy) String() string {
	return string(t)
}
