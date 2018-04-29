package bindings

import "errors"

type Validatable interface {
	Validate() error
}

var ErrNotValidatable = errors.New("Type is not validatable")

type Validator struct{}

func (v *Validator) Validate(i interface{}) error {
	if validatable, ok := i.(Validatable); ok {
		return validatable.Validate()
	}
	return ErrNotValidatable
}
