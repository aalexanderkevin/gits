package model

import (
	"test3/helper"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Publisher struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

func (p Publisher) Validate() (err error) {
	if err = validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
	); err != nil {
		return helper.NewParameterError(helper.Pointer(err.Error()))
	}
	return nil
}
