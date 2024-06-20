package rest

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

var vd *validator.Validate

// Validate validates the incoming request
func Validate(i interface{}) error {
	return vd.Struct(i)
}

func NewValidator() error {
	vd = validator.New()
	if vd == nil {
		return errors.New("failed to initialize validator")
	}

	//vd.RegisterCustomTypeFunc(ValidateValuer, null.String{})
	return nil
}

// FieldValidationFunc field validation function
type FieldValidationFunc func(parent reflect.Value, val reflect.Value, param string) bool

//ValidateValuer implements validator.CustomTypeFunc
//func ValidateValuer(field reflect.Value) interface{} {
//	if str, ok := field.Interface().(null.String); ok {
//		if str.Valid == false {
//			return false
//		}
//	}
//
//	return true
//}
