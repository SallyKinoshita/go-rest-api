package cuuid

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	testUUID = "f47ac10b-58cc-0372-8567-0e02b2c3d479"
)

// uuidライブラリをラップしているだけなので、失敗のテストケースは省略
func TestNewCUUID(t *testing.T) {
	t.Parallel()
	successTests := []struct {
		name string
		want CUUID
	}{
		{
			name: "success",
			want: CUUID{},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := New()
			if err != nil {
				t.Errorf("NewCUUID() error = %v, wantErr %v", err, nil)
				return
			}

			if got.Validate() != nil {
				t.Errorf("NewCUUID() returned an invalid UUID: %v", err)
			}
		})
	}
}

func TestUUID_Validate(t *testing.T) {
	t.Parallel()
	successTests := []struct {
		name string
		u    CUUID
	}{
		{
			name: "success",
			u:    CUUID{testUUID},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.u.Validate(); err != nil {
				t.Errorf("UUID.Validate() error = %v", err)
			}
		})
	}

	failureTests := []struct {
		name string
		u    CUUID
	}{
		{
			name: "failure: uuid is empty",
			u:    CUUID{},
		},
		{
			name: "failure: uuid is invalid",
			u:    CUUID{"invalid-uuid"},
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.u.Validate(); err == nil {
				t.Errorf("UUID.Validate() error = %v, want error = %v", err, "is not nil")
			}
		})
	}
}

func TestUUID_String(t *testing.T) {
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
				string: testUUID,
			},
			want: testUUID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := CUUID{
				string: tt.fields.string,
			}
			got := u.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("UUID.String() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestConvertCUUID(t *testing.T) {
	t.Parallel()
	type args struct {
		uuid string
	}

	successTests := []struct {
		name string
		args args
		want CUUID
	}{
		{
			name: "success",
			args: args{uuid: testUUID},
			want: CUUID{testUUID},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Convert(tt.args.uuid)
			if err != nil {
				t.Errorf("ConvertCUUID() error = %v, wantErr %v", err, nil)
				return
			}

			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("ConvertCUUID() mismatch (-got +want):\n%s", diff)
			}
		})
	}

	failureTests := []struct {
		name string
		args args
		want CUUID
	}{
		{
			name: "failure: empty",
			args: args{uuid: ""},
			want: CUUID{},
		},
		{
			name: "failure: invalid uuid",
			args: args{uuid: "invalid-uuid"},
			want: CUUID{},
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Convert(tt.args.uuid)
			if err == nil {
				t.Errorf("ConvertCUUID() error = %v, wantErr %v", err, "is not nil")
				return
			}

			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("ConvertCUUID() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_TestCUUID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want CUUID
	}{
		{
			name: "success",
			want: CUUID{"9f778116-7c3a-404e-abab-29e0d535dea6"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := TestCUUID()
			opt := cmp.AllowUnexported(got)
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("TestCUUID() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
