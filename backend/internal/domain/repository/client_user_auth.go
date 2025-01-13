package repository

import (
	"context"
	"time"

	"github.com/SallyKinoshita/go-rest-api/backend/internal/domain/model"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cuuid"
)

type ClientUserAuth interface {
	PreSignUp(ctx context.Context, clientUserID cuuid.CUUID, emailAddress model.EmailAddress, password string) (cognitoUserID string, err error)
	ResendConfirmationCode(ctx context.Context, clientUserID cuuid.CUUID) error
	SignUp(ctx context.Context, clientUserID cuuid.CUUID, code string) error
	SignIn(ctx context.Context, clientUserID cuuid.CUUID, password string) (*model.AuthToken, error)
	UpdateEmailAddress(ctx context.Context, clientUserID cuuid.CUUID, email model.EmailAddress) error
	RevokeToken(ctx context.Context, refreshToken string) error
	TokenRefresh(ctx context.Context, clientUserID cuuid.CUUID, refreshToken string) (*model.AuthToken, error)
	Delete(ctx context.Context, clientUserID cuuid.CUUID) error
	GetCognitoUserIDByIDTokenAndValidate(ctx context.Context, idToken string, now time.Time, checkTokenExpiresAt bool) (string, error)
}
