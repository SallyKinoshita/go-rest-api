package cuuid

import (
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	"github.com/google/uuid"
)

type CUUID struct {
	string
}

func New() (CUUID, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return CUUID{}, cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	u := CUUID{uuid.String()}

	if err := u.Validate(); err != nil {
		return CUUID{}, err
	}

	return u, nil
}

func (u CUUID) Validate() error {
	if _, err := uuid.Parse(u.string); err != nil {
		return cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	return nil
}

func (u CUUID) String() string {
	return u.string
}

func Convert(uuid string) (CUUID, error) {
	u := CUUID{uuid}

	if err := u.Validate(); err != nil {
		return CUUID{}, err
	}

	return u, nil
}

func TestCUUID() CUUID {
	return CUUID{"9f778116-7c3a-404e-abab-29e0d535dea6"}
}
