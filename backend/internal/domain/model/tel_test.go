package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_newTel(t *testing.T) {
	t.Parallel()
	type args struct {
		tel string
	}

	successTests := []struct {
		name string
		args args
		want Tel
	}{
		{
			name: "success",
			args: args{
				tel: "+819012345678",
			},
			want: Tel{
				string: "+819012345678",
			},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := newTel(tt.args.tel)
			if err != nil {
				t.Errorf("newTel() error = %v, want nil", err)
				return
			}

			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("newTel() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestTel_validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		string string
	}

	successTests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				string: "+819012345678",
			},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := Tel{
				string: tt.fields.string,
			}

			if err := tr.validate(); err != nil {
				t.Errorf("Tel.validate() error = %v, want nil", err)
			}
		})
	}

	failureTests := []struct {
		name   string
		fields fields
	}{
		{
			name: "failure: tel is empty",
			fields: fields{
				string: "",
			},
		},
		{
			name: "failure: tel is invalid",
			fields: fields{
				string: "invalid",
			},
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := Tel{
				string: tt.fields.string,
			}

			if err := tr.validate(); err == nil {
				t.Errorf("Tel.validate() error = %v, wantErr = is not nil", err)
			}
		})
	}
}

func TestTel_String(t *testing.T) {
	t.Parallel()
	type fields struct {
		string string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				string: "+819012345678",
			},
			want: "+819012345678",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := Tel{
				string: tt.fields.string,
			}

			got := tr.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Tel.String() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
