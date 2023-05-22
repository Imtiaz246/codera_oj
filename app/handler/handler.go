package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/imtiaz246/codera_oj/app/store"
)

type Handler struct {
	*store.Store
	validator *Validator
}

type Validator struct {
	validator *validator.Validate
}

func NewHandler() (*Handler, error) {
	store, err := store.NewStore()
	if err != nil {
		return nil, err
	}
	validator := NewValidator()
	return &Handler{
		store,
		validator,
	}, nil
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate validates a user defined structs
func (v *Validator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err != nil {
		var errMsg string
		for _, e := range err.(validator.ValidationErrors) {
			errMsg += fmt.Sprintf("%s field validation failed on tag '%s', actual value is '%s'\n",
				e.Field(), e.Tag(), e.Value())
		}
		return fmt.Errorf(errMsg)
	}

	return nil
}
