package model

import (
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	nameMaxLength = 50
	nameMinLength = 1
)

type Name struct {
	first     string
	last      string
	firstKana string
	lastKana  string
}

func newName(first, last, firstKana, lastKana string) (Name, error) {
	n := Name{
		first:     first,
		last:      last,
		firstKana: firstKana,
		lastKana:  lastKana,
	}

	if err := n.validate(); err != nil {
		return Name{}, err
	}

	return n, nil
}

func (n Name) validate() error {
	if err := validation.ValidateStruct(&n,
		validation.Field(&n.first, validation.Required, validation.Length(nameMinLength, nameMaxLength)),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailUnprocessableEntityFirstName)
	}

	if err := validation.ValidateStruct(&n,
		validation.Field(&n.last, validation.Required, validation.Length(nameMinLength, nameMaxLength)),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailUnprocessableEntityLastName)
	}

	if err := validation.ValidateStruct(&n,
		validation.Field(&n.firstKana, validation.Required, validation.Length(nameMinLength, nameMaxLength), validation.Match(regKanaOnly)),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailUnprocessableEntityFirstNameKana)
	}

	if err := validation.ValidateStruct(&n,
		validation.Field(&n.lastKana, validation.Required, validation.Length(nameMinLength, nameMaxLength), validation.Match(regKanaOnly)),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailUnprocessableEntityLastNameKana)
	}

	return nil
}

func (n Name) Full() string {
	return n.last + n.first
}

func (n Name) FullKana() string {
	return n.lastKana + n.firstKana
}

func (n Name) First() string {
	return n.first
}

func (n Name) Last() string {
	return n.last
}

func (n Name) FirstKana() string {
	return n.firstKana
}

func (n Name) LastKana() string {
	return n.lastKana
}
