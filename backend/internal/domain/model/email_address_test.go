package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testEmailAddress = "sample@example.com"
)

func TestNewEmailAddress(t *testing.T) {
	t.Parallel()
	type args struct {
		emailAddress string
	}

	successTests := []struct {
		name string
		args args
		want EmailAddress
	}{
		{
			name: "success",
			args: args{
				emailAddress: testEmailAddress,
			},
			want: EmailAddress{testEmailAddress},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewEmailAddress(tt.args.emailAddress)
			if err != nil {
				t.Errorf("newEmailAddress() error = %v, wantErr %v", err, nil)
				return
			}

			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("newEmailAddress() mismatch (-got +want):\n%s", diff)
			}
		})
	}

	failureTests := []struct {
		name string
		args args
		want EmailAddress
	}{
		{
			name: "failure: email address is empty",
			args: args{
				emailAddress: "",
			},
			want: EmailAddress{},
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewEmailAddress(tt.args.emailAddress)
			if err == nil {
				t.Errorf("newEmailAddress() error = %v, wantErr = is not nil", err)
				return
			}

			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("newEmailAddress() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestEmailAddress_validate(t *testing.T) {
	t.Parallel()

	successTests := []struct {
		name string
		e    EmailAddress
	}{
		{
			name: "success",
			e:    EmailAddress{testEmailAddress},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.e.validate(); err != nil {
				t.Errorf("EmailAddress.validate() error = %v, wantErr = is nil", err)
			}
		})
	}

	failureTests := []struct {
		name string
		e    EmailAddress
	}{
		{
			name: "failure: email address is empty",
			e:    EmailAddress{},
		},
		{
			name: "failure: user name is empty",
			e:    EmailAddress{"@example.com"},
		},
		{
			name: "failure: domain is empty",
			e:    EmailAddress{"sample@"},
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.e.validate(); err == nil {
				t.Errorf("EmailAddress.validate() error = %v, wantErr = is not nil", err)
			}
		})
	}
}

func TestEmailAddress_String(t *testing.T) {
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
				string: testEmailAddress,
			},
			want: testEmailAddress,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := EmailAddress{
				string: tt.fields.string,
			}
			got := e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("EmailAddress.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
