package model

import (
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Tel struct {
	string
}

func newTel(tel string) (Tel, error) {
	t := Tel{tel}

	if err := t.validate(); err != nil {
		return Tel{}, err
	}

	return t, nil
}

func (t Tel) validate() error {
	if err := validation.ValidateStruct(&t,
		validation.Field(&t.string, validation.Required, is.E164),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailUnprocessableEntityTel)
	}

	return nil
}

func (t Tel) String() string {
	return t.string
}
