package model

import (
	"test3/helper"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Author struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

func (a Author) Validate() (err error) {
	if err = validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
	); err != nil {
		return helper.NewParameterError(helper.Pointer(err.Error()))
	}
	return nil
}
