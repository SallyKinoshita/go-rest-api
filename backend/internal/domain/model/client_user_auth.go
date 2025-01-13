package model

import (
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AuthToken struct {
	idToken      string
	accessToken  string
	refreshToken string
}

func NewAuthToken(idToken, accessToken, refreshToken string) (*AuthToken, error) {
	authToken := &AuthToken{
		idToken:      idToken,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}

	if err := authToken.validate(); err != nil {
		return nil, err
	}

	return authToken, nil
}

func (a AuthToken) validate() error {
	if err := validation.ValidateStruct(
		&a,
		validation.Field(&a.idToken, validation.Required),
		validation.Field(&a.accessToken, validation.Required),
		validation.Field(&a.refreshToken, validation.Required),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	return nil
}

func (a AuthToken) IDToken() string {
	return a.idToken
}

func (a AuthToken) AccessToken() string {
	return a.accessToken
}

func (a AuthToken) RefreshToken() string {
	return a.refreshToken
}
