package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewPagination(t *testing.T) {
	t.Parallel()
	type args struct {
		page  int
		limit int
	}

	tests := []struct {
		name string
		args args
		want Pagination
	}{
		{
			name: "success",
			args: args{
				page:  2,
				limit: 10,
			},
			want: Pagination{
				offset: 10,
				limit:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewPagination(tt.args.page, tt.args.limit)
			if err != nil {
				t.Errorf("NewPagination() error = %v, wantErr %v", err, nil)
				return
			}

			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("NewPagination() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_calOffset(t *testing.T) {
	t.Parallel()

	type args struct {
		page  int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				page:  2,
				limit: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := calOffset(tt.args.page, tt.args.limit)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("calOffset() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestPagination_validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		offset int
		limit  int
	}

	successTests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				offset: 1,
				limit:  10,
			},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := Pagination{
				offset: tt.fields.offset,
				limit:  tt.fields.limit,
			}
			if err := p.validate(); err != nil {
				t.Errorf("validate() error = %v, wantErr %v", err, nil)
			}
		})
	}

	failureTests := []struct {
		name   string
		fields fields
	}{
		{
			name: "failure: offset is less than offset min",
			fields: fields{
				offset: 0,
				limit:  10,
			},
		},
		{
			name: "failure: limit is less than limit min",
			fields: fields{
				offset: 1,
				limit:  0,
			},
		},
		{
			name: "failure: limit is greater than limit max",
			fields: fields{
				offset: 1,
				limit:  101,
			},
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := Pagination{
				offset: tt.fields.offset,
				limit:  tt.fields.limit,
			}
			if err := p.validate(); err == nil {
				t.Errorf("validate() error = %v, wantErr = is not nil", err)
			}
		})
	}
}

func TestPagination_Offset(t *testing.T) {
	t.Parallel()
	type fields struct {
		offset int
		limit  int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "success",
			fields: fields{
				offset: 1,
				limit:  10,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := Pagination{
				offset: tt.fields.offset,
				limit:  tt.fields.limit,
			}

			got := p.Offset()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Offset() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestPagination_Limit(t *testing.T) {
	t.Parallel()
	type fields struct {
		offset int
		limit  int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "success",
			fields: fields{
				offset: 1,
				limit:  10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := Pagination{
				offset: tt.fields.offset,
				limit:  tt.fields.limit,
			}

			got := p.Limit()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Limit() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
