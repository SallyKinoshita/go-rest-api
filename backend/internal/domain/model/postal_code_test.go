package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_newPostalCode(t *testing.T) {
	t.Parallel()
	type args struct {
		pc string
	}

	tests := []struct {
		name string
		args args
		want postalCode
	}{
		{
			name: "success",
			args: args{
				pc: "1234567",
			},
			want: postalCode{"1234567"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := newPostalCode(tt.args.pc)
			if err != nil {
				t.Errorf("newPostalCode() error = %v, wantErr %v", err, nil)
				return
			}

			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("newPostalCode() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_postalCode_validate(t *testing.T) {
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
				string: "1234567",
			},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := postalCode{
				string: tt.fields.string,
			}

			if err := p.validate(); err != nil {
				t.Errorf("postalCode.validate() error = %v, wantErr %v", err, nil)
			}
		})
	}

	failureTests := []struct {
		name   string
		fields fields
	}{
		{
			name: "failure: post code is empty",
			fields: fields{
				string: "",
			},
		},
		{
			name: "failure: post code is not 7 digits",
			fields: fields{
				string: "123456",
			},
		},
		{
			name: "failure: post code is not digit",
			fields: fields{
				string: "123456a",
			},
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := postalCode{
				string: tt.fields.string,
			}

			if err := p.validate(); err == nil {
				t.Errorf("postalCode.validate() error = %v, wantErr = is not nil", err)
			}
		})
	}
}

func Test_postalCode_Format(t *testing.T) {
	t.Parallel()
	type fields struct {
		string string
	}

	tests := []struct {
		name   string
		fields fields
		want   postalCode
	}{
		{
			name: "success",
			fields: fields{
				string: "1234567",
			},
			want: postalCode{"123-4567"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := postalCode{
				string: tt.fields.string,
			}

			got := p.Format()
			opt := cmp.AllowUnexported(got)

			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("postalCode.Format() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_postalCode_String(t *testing.T) {
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
				string: "123-4567",
			},
			want: "123-4567",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := postalCode{
				string: tt.fields.string,
			}

			got := p.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("postalCode.String() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
