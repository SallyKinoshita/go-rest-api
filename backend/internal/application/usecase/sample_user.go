package usecase

import (
	"context"
	"time"

	"github.com/SallyKinoshita/go-rest-api/backend/internal/domain/model"
	"github.com/SallyKinoshita/go-rest-api/backend/internal/domain/repository"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/ccontext"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cuuid"
	"github.com/uptrace/bun"
)

type SampleUserUseCase interface {
	Create(ctx context.Context, firstName, lastName, firstNameKana, lastNameKana, emailAddress, tel string, birthDate time.Time, postalCode string, prefecture model.Prefecture, city, streetAndNumber, buildingAndRoom string) error
	Update(ctx context.Context, sampleUserID, firstName, lastName, firstNameKana, lastNameKana, emailAddress, postalCode string, prefecture model.Prefecture, city, streetAndNumber, buildingAndRoom string) error
	Get(ctx context.Context, sampleUserID cuuid.CUUID) (*model.SampleUser, error)
	List(ctx context.Context, page, limit int) (model.SampleUsers, error)
	Delete(ctx context.Context, sampleUserID string) error
}

type SampleUser struct {
	db             *bun.DB
	sampleUserRepo repository.SampleUser
}

func NewSampleUser(db *bun.DB, sampleUserRepo repository.SampleUser) *SampleUser {
	return &SampleUser{
		db:             db,
		sampleUserRepo: sampleUserRepo,
	}
}

// TODO: テスト実装する (本ファイルの全関数)

func (s *SampleUser) Create(ctx context.Context, firstName, lastName, firstNameKana, lastNameKana, emailAddress, tel string, birthDate time.Time, postalCode string, prefecture model.Prefecture, city, streetAndNumber, buildingAndRoom string) error {
	now, err := ccontext.GetNowTime(ctx)
	if err != nil {
		return err
	}

	sampleUserID, err := cuuid.New()
	if err != nil {
		return err
	}

	user, err := model.NewSampleUser(sampleUserID, firstName, lastName, firstNameKana, lastNameKana, emailAddress, tel, birthDate, postalCode, prefecture, city, streetAndNumber, buildingAndRoom, now)
	if err != nil {
		return err
	}

	return s.sampleUserRepo.Create(ctx, s.db, now, user)
}

func (s *SampleUser) Update(ctx context.Context, sampleUserID, firstName, lastName, firstNameKana, lastNameKana, emailAddress, postalCode string, prefecture model.Prefecture, city, streetAndNumber, buildingAndRoom string) error {
	now, err := ccontext.GetNowTime(ctx)
	if err != nil {
		return err
	}
	uuid, err := cuuid.Convert(sampleUserID)
	if err != nil {
		return err
	}

	user, err := s.sampleUserRepo.Get(ctx, s.db, uuid)
	if err != nil {
		return err
	}

	if err := user.Update(firstName, lastName, firstNameKana, lastNameKana, emailAddress, postalCode, prefecture, city, streetAndNumber, buildingAndRoom); err != nil {
		return err
	}

	return s.sampleUserRepo.Update(ctx, s.db, now, user)
}

func (s *SampleUser) Get(ctx context.Context, sampleUserID cuuid.CUUID) (*model.SampleUser, error) {
	return s.sampleUserRepo.Get(ctx, s.db, sampleUserID)
}

func (s *SampleUser) List(ctx context.Context, page, limit int) (model.SampleUsers, error) {
	pagination, err := model.NewPagination(page, limit)
	if err != nil {
		return nil, err
	}

	return s.sampleUserRepo.List(ctx, s.db, pagination)
}

func (s *SampleUser) Delete(ctx context.Context, sampleUserID string) error {
	uuid, err := cuuid.Convert(sampleUserID)
	if err != nil {
		return err
	}

	return s.sampleUserRepo.Delete(ctx, s.db, uuid)
}
