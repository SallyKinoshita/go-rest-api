package middleware

import (
	"context"
	"strings"

	"github.com/SallyKinoshita/go-rest-api/backend/internal/domain/repository"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/ccontext"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type ClientUser struct {
	db                 *bun.DB
	clientUserRepo     repository.SampleUser
	clientUserAuthRepo repository.ClientUserAuth
}

func newClientUser(db *bun.DB, clientUserRepo repository.SampleUser, clientUserAuthRepo repository.ClientUserAuth) *ClientUser {
	return &ClientUser{
		db:                 db,
		clientUserRepo:     clientUserRepo,
		clientUserAuthRepo: clientUserAuthRepo,
	}
}

func (c *ClientUser) CognitoAuth(ctx context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			if checkPath(ec.Path(), clientUserSkipAuthPath) {
				return next(ec)
			}

			authHeader := ec.Request().Header.Get(echo.HeaderAuthorization)

			idToken := strings.TrimPrefix(authHeader, "Bearer ")
			if idToken == "" {
				return cerror.New("invalid auth header", cerror.DetailUnauthorized)
			}

			var checkTokenExpiresAt bool
			if ec.Path() != sampleUserTokenRefreshPath {
				checkTokenExpiresAt = true
			}

			ctx := ec.Request().Context()

			now, err := ccontext.GetNowTime(ctx)
			if err != nil {
				return err
			}

			sub, err := c.clientUserAuthRepo.GetCognitoUserIDByIDTokenAndValidate(ctx, idToken, now, checkTokenExpiresAt)
			if err != nil {
				return err
			}

			clientUserID, err := c.clientUserRepo.GetIDByCognitoUserID(ctx, c.db, sub)
			if err != nil {
				return err
			}

			ctx = ccontext.SetClientUserID(ctx, clientUserID)

			ec.SetRequest(ec.Request().WithContext(ctx))

			return next(ec)
		}
	}
}
