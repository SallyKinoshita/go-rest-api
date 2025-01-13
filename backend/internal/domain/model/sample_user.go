package model

import (
	"time"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cuuid"
)

type SampleUsers []*SampleUser

type SampleUser struct {
	id           cuuid.CUUID
	name         Name
	emailAddress EmailAddress
	tel          Tel
	birthDate    BirthDate
	address      Address
}

func NewSampleUser(sampleUserID cuuid.CUUID, firstName, lastName, firstNameKana, lastNameKana, emailAddress, tel string, birthDate time.Time, postalCode string, prefecture Prefecture, city, streetAndNumber, buildingAndRoom string, now time.Time) (*SampleUser, error) {
	if err := sampleUserID.Validate(); err != nil {
		return nil, err
	}

	name, err := newName(firstName, lastName, firstNameKana, lastNameKana)
	if err != nil {
		return nil, err
	}

	e, err := NewEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	t, err := newTel(tel)
	if err != nil {
		return nil, err
	}

	b, err := newBirthDate(birthDate, now)
	if err != nil {
		return nil, err
	}

	a, err := newAddress(postalCode, prefecture, city, streetAndNumber, buildingAndRoom)
	if err != nil {
		return nil, err
	}

	return &SampleUser{
		id:           sampleUserID,
		name:         name,
		emailAddress: e,
		tel:          t,
		birthDate:    b,
		address:      *a,
	}, nil
}

func NewSampleUserFromDB(sampleUserID cuuid.CUUID, firstName, lastName, firstNameKana, lastNameKana, emailAddress, tel string, birthDate time.Time, postalCode string, prefecture string, city, streetAndNumber, buildingAndRoom string, createdAt, updatedAt time.Time) (*SampleUser, error) {
	if err := sampleUserID.Validate(); err != nil {
		return nil, err
	}

	name, err := newName(firstName, lastName, firstNameKana, lastNameKana)
	if err != nil {
		return nil, err
	}

	e, err := NewEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	t, err := newTel(tel)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	b, err := newBirthDate(birthDate, now)
	if err != nil {
		return nil, err
	}

	p, err := PrefectureKeyByValue(prefecture)
	if err != nil {
		return nil, err
	}

	a, err := newAddress(postalCode, p, city, streetAndNumber, buildingAndRoom)
	if err != nil {
		return nil, err
	}

	return &SampleUser{
		id:           sampleUserID,
		name:         name,
		emailAddress: e,
		tel:          t,
		birthDate:    b,
		address:      *a,
	}, nil
}

func (s *SampleUser) Update(firstName, lastName, firstNameKana, lastNameKana, emailAddress, postalCode string, prefecture Prefecture, city, streetAndNumber, buildingAndRoom string) error {
	name, err := newName(firstName, lastName, firstNameKana, lastNameKana)
	if err != nil {
		return err
	}

	e, err := NewEmailAddress(emailAddress)
	if err != nil {
		return err
	}

	a, err := newAddress(postalCode, prefecture, city, streetAndNumber, buildingAndRoom)
	if err != nil {
		return err
	}

	s.name = name
	s.emailAddress = e
	s.address = *a

	return nil
}

func (s SampleUser) ID() cuuid.CUUID {
	return s.id
}

func (s SampleUser) Name() Name {
	return s.name
}

func (s SampleUser) EmailAddress() EmailAddress {
	return s.emailAddress
}

func (s SampleUser) Tel() Tel {
	return s.tel
}

func (s SampleUser) BirthDate() BirthDate {
	return s.birthDate
}

func (s SampleUser) Address() Address {
	return s.address
}
