package presenter

import (
	"github.com/SallyKinoshita/go-rest-api/backend/internal/domain/model"
	"github.com/SallyKinoshita/go-rest-api/backend/internal/gen/openapi"
	"github.com/labstack/echo/v4"
)

func toPtr(s string) *string {
	return &s
}

func NewGetSampleUserResponse(ec echo.Context, user *model.SampleUser) *openapi.SampleServiceGetSampleUserResponse {
	return &openapi.SampleServiceGetSampleUserResponse{
		JSON200: &openapi.V1GetSampleUserResponse{
			//TODO: ちゃんと詰め直す
			SampleUserId: toPtr(user.ID().String()),
			// FirstName:       toPtr(user.Name().First),
			// LastName:        toPtr(user.Name().Last),
			// FirstNameKana:   toPtr(user.Name().FirstKana),
			// LastNameKana:    toPtr(user.Name().LastKana),
			EmailAddress: toPtr(user.EmailAddress().String()),
			Tel:          toPtr(user.Tel().String()),
			// BirthDate:       toPtr(user.BirthDate().String()),
			// Prefecture:      toPtr(user.Address().Prefecture.String()),
			// City:            toPtr(user.Address().City),
			// StreetAndNumber: toPtr(user.Address().StreetAndNumber),
			// BuildingAndRoom: toPtr(user.Address().BuildingAndRoom),
		},
	}
}

func NewListSampleUserResponse(ec echo.Context, users model.SampleUsers) *openapi.SampleServiceListSampleUserResponse {
	res := make([]openapi.V1ListSampleUserRes, len(users))

	for i, v := range users {
		res[i] = openapi.V1ListSampleUserRes{
			//TODO: ちゃんと詰め直す
			SampleUserId: toPtr(v.ID().String()),
			// FirstName:       toPtr(v.Name().First),
			// LastName:        toPtr(v.Name().Last),
			// FirstNameKana:   toPtr(v.Name().FirstKana),
			// LastNameKana:    toPtr(v.Name().LastKana),
			EmailAddress: toPtr(v.EmailAddress().String()),
			Tel:          toPtr(v.Tel().String()),
			// BirthDate:       toPtr(v.BirthDate().String()),
			// Prefecture:      toPtr(v.Address().Prefecture.String()),
			// City:            toPtr(v.Address().City),
			// StreetAndNumber: toPtr(v.Address().StreetAndNumber),
			// BuildingAndRoom: toPtr(v.Address().BuildingAndRoom),
		}
	}

	return &openapi.SampleServiceListSampleUserResponse{
		JSON200: &openapi.V1ListSampleUserResponse{
			SampleUsers: &res,
		},
	}
}
