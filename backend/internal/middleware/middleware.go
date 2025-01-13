package middleware

import (
	"github.com/SallyKinoshita/go-rest-api/backend/internal/domain/repository"
	"github.com/uptrace/bun"
)

type Middleware struct {
	AccessLog  *AccessLog
	CLog       *CLog
	CError     *CError
	ClientUser *ClientUser
}

func New(
	db *bun.DB,
	sampleUserRepo repository.SampleUser,
	clientUserAuthRepo repository.ClientUserAuth,
) *Middleware {
	return &Middleware{
		AccessLog:  newAccessLog(),
		CLog:       newCLog(),
		CError:     newCError(),
		ClientUser: newClientUser(db, sampleUserRepo, clientUserAuthRepo),
	}
}
