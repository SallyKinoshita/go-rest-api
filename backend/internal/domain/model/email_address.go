package model

import (
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type EmailAddress struct {
	string
}

func NewEmailAddress(emailAddress string) (EmailAddress, error) {
	e := EmailAddress{emailAddress}

	if err := e.validate(); err != nil {
		return EmailAddress{}, err
	}

	return e, nil
}

func (e EmailAddress) validate() error {
	if err := validation.ValidateStruct(&e,
		validation.Field(&e.string, validation.Required, is.Email),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailUnprocessableEntityEmailAddress)
	}

	return nil
}

func (e EmailAddress) String() string {
	return e.string
}
