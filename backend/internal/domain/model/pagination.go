package model

import (
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	limitMin = 1
	limitMax = 100
)

type Pagination struct {
	offset int
	limit  int
}

func NewPagination(page, limit int) (Pagination, error) {
	offset := calOffset(page, limit)

	p := Pagination{
		offset: offset,
		limit:  limit,
	}

	if err := p.validate(); err != nil {
		return Pagination{}, err
	}

	return p, nil
}

func calOffset(page, limit int) int {
	return limit * (page - 1)
}

func (p Pagination) validate() error {
	if err := validation.ValidateStruct(&p,
		validation.Field(&p.offset, validation.Required),
		validation.Field(&p.limit, validation.Required, validation.Min(limitMin), validation.Max(limitMax)),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailBadRequest)
	}

	return nil
}

func (p Pagination) Offset() int {
	return p.offset
}

func (p Pagination) Limit() int {
	return p.limit
}
