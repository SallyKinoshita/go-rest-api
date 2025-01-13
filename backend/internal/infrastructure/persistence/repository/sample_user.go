package persistencerepository

import (
	"context"
	"database/sql"
	"time"

	"github.com/SallyKinoshita/go-rest-api/backend/internal/domain/model"
	"github.com/SallyKinoshita/go-rest-api/backend/internal/domain/repository"
	persistencemodel "github.com/SallyKinoshita/go-rest-api/backend/internal/infrastructure/persistence/model"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cuuid"
)

type SampleUser struct {
	dbSampleUser persistencemodel.SampleUser
}

func NewSampleUser() *SampleUser {
	return &SampleUser{}
}

func (s *SampleUser) Create(ctx context.Context, conn repository.DBConn, now time.Time, clientUser *model.SampleUser) error {
	cu := persistencemodel.NewSampleUser(now, *clientUser)

	result, err := conn.NewInsert().Model(cu).Exec(ctx)
	if err != nil {
		return cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	if rows == 0 {
		return cerror.New("failed to create client user", cerror.DetailInternalServerError)
	}

	return nil
}

func (s *SampleUser) Update(ctx context.Context, conn repository.DBConn, now time.Time, clientUser *model.SampleUser) error {
	cu := persistencemodel.NewSampleUser(now, *clientUser)

	// created_atは更新しない
	result, err := conn.NewUpdate().Model(cu).ExcludeColumn("created_at").Where("client_user_id = ?", cu.ID).Exec(ctx)
	if err != nil {
		return cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	if rows == 0 {
		return cerror.New("failed to update client user", cerror.DetailInternalServerError)
	}

	return nil
}

func (s *SampleUser) Get(ctx context.Context, conn repository.DBConn, clientUserID cuuid.CUUID) (*model.SampleUser, error) {
	if err := conn.NewSelect().Model(&s.dbSampleUser).Where("client_user_id = ?", clientUserID.String()).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, cerror.Wrap(err, cerror.DetailNotFoundClientUser)
		}
		return nil, cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	cu, err := cuuid.Convert(s.dbSampleUser.ID)
	if err != nil {
		return nil, err
	}

	clientUser, err := model.NewSampleUserFromDB(cu, s.dbSampleUser.FirstName, s.dbSampleUser.LastName, s.dbSampleUser.FirstNameKana, s.dbSampleUser.LastNameKana, s.dbSampleUser.EmailAddress, s.dbSampleUser.Tel, s.dbSampleUser.BirthDate, s.dbSampleUser.PostalCode, s.dbSampleUser.Prefecture, s.dbSampleUser.City, s.dbSampleUser.StreetAndNumber, s.dbSampleUser.BuildingAndRoom, s.dbSampleUser.CreatedAt, s.dbSampleUser.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return clientUser, nil
}

func (s *SampleUser) GetIDByCognitoUserID(ctx context.Context, conn repository.DBConn, cognitoUserID string) (cuuid.CUUID, error) {
	if err := conn.NewSelect().Model(&s.dbSampleUser).Where("cognito_user_id = ?", cognitoUserID).Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return cuuid.CUUID{}, cerror.Wrap(err, cerror.DetailNotFoundClientUser)
		}
		return cuuid.CUUID{}, cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	id, err := cuuid.Convert(s.dbSampleUser.ID)
	if err != nil {
		return cuuid.CUUID{}, err
	}

	return id, nil
}

func (s *SampleUser) Delete(ctx context.Context, conn repository.DBConn, clientUserID cuuid.CUUID) error {
	result, err := conn.NewDelete().Model(&s.dbSampleUser).Where("client_user_id = ?", clientUserID.String()).Exec(ctx)
	if err != nil {
		return cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	if rows == 0 {
		return cerror.New("failed to delete client user", cerror.DetailNotFoundClientUser)
	}

	return nil
}

func (s *SampleUser) List(ctx context.Context, conn repository.DBConn, pagination model.Pagination) (model.SampleUsers, error) {
	var sampleUsers []persistencemodel.SampleUser

	query := conn.NewSelect().Model(&sampleUsers)

	// ページネーションの適用
	if pagination.Limit() > 0 {
		query = query.Limit(pagination.Limit())
	}
	if pagination.Offset() > 0 {
		query = query.Offset(pagination.Offset())
	}

	if err := query.Scan(ctx); err != nil {
		if err == sql.ErrNoRows {
			return nil, cerror.Wrap(err, cerror.DetailNotFoundClientUser)
		}
		return nil, cerror.Wrap(err, cerror.DetailInternalServerError)
	}

	var result model.SampleUsers
	for _, su := range sampleUsers {
		uuid, err := cuuid.Convert(su.ID)
		if err != nil {
			return nil, err
		}

		domainSampleUser, err := model.NewSampleUserFromDB(
			uuid,
			su.FirstName,
			su.LastName,
			su.FirstNameKana,
			su.LastNameKana,
			su.EmailAddress,
			su.Tel,
			su.BirthDate,
			su.PostalCode,
			su.Prefecture,
			su.City,
			su.StreetAndNumber,
			su.BuildingAndRoom,
			su.CreatedAt,
			su.UpdatedAt,
		)
		if err != nil {
			return nil, cerror.Wrap(err, cerror.DetailInternalServerError)
		}
		result = append(result, domainSampleUser)
	}

	return result, nil
}
