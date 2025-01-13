package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/SallyKinoshita/go-rest-api/backend/internal/domain/model"
	"github.com/SallyKinoshita/go-rest-api/backend/internal/gen/openapi"
	"github.com/SallyKinoshita/go-rest-api/backend/internal/interfaces/presenter"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/ccontext"
	"github.com/labstack/echo/v4"
)

type SampleUserController struct {
	sampleUserUseCase SampleUserUseCase
}

type SampleUserUseCase interface {
	Create(ctx context.Context, firstName, lastName, firstNameKana, lastNameKana, emailAddress, tel string, birthDate time.Time, postalCode string, prefecture int, city, streetAndNumber, buildingAndRoom string) error
	Update(ctx context.Context, sampleUserID, firstName, lastName, firstNameKana, lastNameKana, emailAddress, postalCode string, prefecture int, city, streetAndNumber, buildingAndRoom string) error
	Get(ctx context.Context, sampleUserID string) (*model.SampleUser, error)
	List(ctx context.Context, page, limit int) (model.SampleUsers, error)
	Delete(ctx context.Context, sampleUserID string) error
}

func (c *Controller) SampleServiceCreateSampleUser(ec echo.Context, params openapi.SampleServiceCreateSampleUserParams) error {
	ctx := ec.Request().Context()

	//TODO: ここでprefectureを変換する
	// prefecture, err := masterenum.ConvertPrefectureModelKey(*params.Prefecture)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Invalid prefecture value")
	// }
	prefecture := model.PrefectureAichi // TODO: 仮の値

	err := c.sampleUserUseCase.Create(ctx, *params.FirstName, *params.LastName, *params.FirstNameKana, *params.LastNameKana, *params.EmailAddress, *params.Tel, *params.BirthDate, *params.PostalCode, prefecture, *params.City, *params.StreetAndNumber, *params.BuildingAndRoom)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create sample user")
	}

	return ec.JSON(http.StatusOK, map[string]any{"message": "User created successfully"})
}

func (c *Controller) SampleServiceUpdateSampleUser(ec echo.Context, params openapi.SampleServiceUpdateSampleUserParams) error {
	ctx := ec.Request().Context()

	//TODO: ここでprefectureを変換する
	// prefecture, err := masterenum.ConvertPrefectureModelKey(*params.Prefecture)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Invalid prefecture value")
	// }
	prefecture := model.PrefectureAichi // TODO: 仮の値

	err := c.sampleUserUseCase.Update(ctx, *params.SampleUserId, *params.FirstName, *params.LastName, *params.FirstNameKana, *params.LastNameKana, *params.EmailAddress, *params.PostalCode, prefecture, *params.City, *params.StreetAndNumber, *params.BuildingAndRoom)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update sample user")
	}

	return ec.JSON(http.StatusOK, map[string]any{"message": "User updated successfully"})
}

func (c *Controller) SampleServiceGetSampleUser(ec echo.Context, params openapi.SampleServiceGetSampleUserParams) error {
	ctx := ec.Request().Context()

	clientUserID, err := ccontext.GetClientUserID(ctx)
	user, err := c.sampleUserUseCase.Get(ctx, clientUserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return ec.JSON(http.StatusOK, presenter.NewGetSampleUserResponse(ec, user))
}

func (c *Controller) SampleServiceListSampleUser(ec echo.Context, params openapi.SampleServiceListSampleUserParams) error {
	ctx := ec.Request().Context()

	users, err := c.sampleUserUseCase.List(ctx, int(*params.Page), int(*params.Limit))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch users")
	}

	return ec.JSON(http.StatusOK, presenter.NewListSampleUserResponse(ec, users))
}

func (c *Controller) SampleServiceDeleteSampleUser(ec echo.Context, params openapi.SampleServiceDeleteSampleUserParams) error {
	ctx := ec.Request().Context()

	err := c.sampleUserUseCase.Delete(ctx, *params.SampleUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
	}

	return ec.JSON(http.StatusOK, map[string]any{"message": "User deleted successfully"})
}
