package model

import (
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	postalCodeLength = 7
)

type postalCode struct {
	string
}

func newPostalCode(pc string) (postalCode, error) {
	p := postalCode{pc}

	if err := p.validate(); err != nil {
		return postalCode{}, err
	}

	return p, nil
}

func (p postalCode) validate() error {
	if err := validation.Validate(p.String(),
		validation.Required, validation.Length(postalCodeLength, postalCodeLength), is.Digit,
	); err != nil {
		return cerror.Wrap(err, cerror.DetailUnprocessableEntityPostalCode)
	}

	return nil
}

func (p postalCode) Format() postalCode {
	f := p.String()[:3] + "-" + p.String()[3:]
	return postalCode{f}
}

func (p postalCode) String() string {
	return p.string
}
