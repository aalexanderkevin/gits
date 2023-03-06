package model

import (
	"test3/helper"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Book struct {
	ID          *string `json:"id"`
	Name        *string `json:"name"`
	AuthorID    *string `json:"author_id"`
	PublisherID *string `json:"publisher_id"`
}

func (b Book) Validate() (err error) {
	if err = validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required),
		validation.Field(&b.AuthorID, validation.Required),
		validation.Field(&b.PublisherID, validation.Required),
	); err != nil {
		return helper.NewParameterError(helper.Pointer(err.Error()))
	}
	return nil
}
