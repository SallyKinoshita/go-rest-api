package model

import (
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	addressMaxLength = 50
	addressMinLength = 1
)

type Address struct {
	postalCode      postalCode // 郵便番号
	prefecture      string     // 都道府県
	city            string     // 市区町村
	streetAndNumber string     // 丁目・番地
	buildingAndRoom string     // 建物名・部屋番号
}

func newAddress(postalCode string, prefecture Prefecture, city, streetAndNumber, buildingAndRoom string) (*Address, error) {
	p, err := prefecture.Value()
	if err != nil {
		return nil, err
	}

	pc, err := newPostalCode(postalCode)
	if err != nil {
		return nil, err
	}

	a := &Address{
		postalCode:      pc,
		prefecture:      p,
		city:            city,
		streetAndNumber: streetAndNumber,
		buildingAndRoom: buildingAndRoom,
	}

	if err := a.validate(); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Address) validate() error {
	if err := validation.ValidateStruct(a,
		validation.Field(&a.city, validation.Required, validation.Length(addressMinLength, addressMaxLength)),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailUnprocessableEntityCity)
	}

	if err := validation.ValidateStruct(a,
		validation.Field(&a.streetAndNumber, validation.Required, validation.Length(addressMinLength, addressMaxLength)),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailUnprocessableEntityStreetAndNumber)
	}

	if err := validation.ValidateStruct(a,
		validation.Field(&a.buildingAndRoom, validation.Length(0, addressMaxLength)),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailUnprocessableEntityBuildingAndRoom)
	}

	return nil
}

func (a Address) PostalCode() postalCode {
	return a.postalCode
}

func (a Address) Prefecture() string {
	return a.prefecture
}

func (a Address) City() string {
	return a.city
}

func (a Address) StreetAndNumber() string {
	return a.streetAndNumber
}

func (a Address) BuildingAndRoom() string {
	return a.buildingAndRoom
}
