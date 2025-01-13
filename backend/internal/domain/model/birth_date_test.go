package model

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var (
	now                = time.Date(2022, time.December, 24, 0, 0, 0, 0, time.Local)
	yesterday          = time.Date(2022, time.December, 23, 0, 0, 0, 0, time.Local)
	february28         = time.Date(2022, time.February, 28, 0, 0, 0, 0, time.Local)
	february27         = time.Date(2022, time.February, 27, 0, 0, 0, 0, time.Local)
	leapYearFebruary29 = time.Date(2020, time.February, 29, 0, 0, 0, 0, time.Local)
	leapYearFebruary28 = time.Date(2020, time.February, 28, 0, 0, 0, 0, time.Local)

	birthDate              = time.Date(1999, time.December, 24, 0, 0, 0, 0, time.Local)
	birthDateLeapYearToDay = time.Date(2000, time.February, 29, 0, 0, 0, 0, time.Local)
)

func Test_newBirthDate(t *testing.T) {
	t.Parallel()

	type args struct {
		birthDate time.Time
		now       time.Time
	}

	tests := []struct {
		name string
		args args
		want BirthDate
	}{
		{
			name: "success",
			args: args{
				birthDate: birthDate,
				now:       now,
			},
			want: BirthDate{Time: birthDate},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := newBirthDate(tt.args.birthDate, tt.args.now)
			if err != nil {
				t.Errorf("newBirthDate() error = %v, wantErr %v", err, nil)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("newBirthDate() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestBirthDate_validate(t *testing.T) {
	t.Parallel()

	type fields struct {
		Time time.Time
	}
	type args struct {
		now time.Time
	}

	successTests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				Time: birthDate,
			},
			args: args{
				now: now,
			},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := BirthDate{
				Time: tt.fields.Time,
			}

			if err := b.validate(tt.args.now); err != nil {
				t.Errorf("BirthDate.validate() error = %v, wantErr %v", err, nil)
			}
		})
	}

	failureTests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "failure: age is less than ageMin",
			fields: fields{
				Time: now,
			},
			args: args{
				now: now,
			},
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := BirthDate{
				Time: tt.fields.Time,
			}

			if err := b.validate(tt.args.now); err == nil {
				t.Error("BirthDate.validate() error = nil, wantErr not nil")
			}
		})
	}
}

func TestBirthDate_CalAge(t *testing.T) {
	t.Parallel()
	type fields struct {
		Time time.Time
	}
	type args struct {
		now time.Time
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "success: age is 23",
			fields: fields{
				Time: birthDate,
			},
			args: args{
				now: now,
			},
			want: 23,
		},
		{
			name: "success: age is 22",
			fields: fields{
				Time: birthDate,
			},
			args: args{
				now: yesterday,
			},
			want: 22,
		},
		{
			name: "success: age is 20 leap year birth on leap year",
			fields: fields{
				Time: birthDateLeapYearToDay,
			},
			args: args{
				now: leapYearFebruary29,
			},
			want: 20,
		},
		{
			name: "success: age is 19 leap year birth on leap year",
			fields: fields{
				Time: birthDateLeapYearToDay,
			},
			args: args{
				now: leapYearFebruary28,
			},
			want: 19,
		},
		{
			name: "success: age is 22 leap year birth not on leap year",
			fields: fields{
				Time: birthDateLeapYearToDay,
			},
			args: args{
				now: february28,
			},
			want: 22,
		},
		{
			name: "success: age is 21 leap year birth not on leap year",
			fields: fields{
				Time: birthDateLeapYearToDay,
			},
			args: args{
				now: february27,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := BirthDate{
				Time: tt.fields.Time,
			}

			if got := b.CalAge(tt.args.now); got != tt.want {
				t.Errorf("BirthDate.CalAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBirthDate_IsLeapYear(t *testing.T) {
	t.Parallel()
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success: leap year",
			fields: fields{
				Time: leapYearFebruary29,
			},
			want: true,
		},
		{
			name: "success: not leap year",
			fields: fields{
				Time: february28,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			b := BirthDate{
				Time: tt.fields.Time,
			}
			got := b.IsLeapYear()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("BirthDate.IsLeapYear() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
