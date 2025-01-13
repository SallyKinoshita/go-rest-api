package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_newAddress(t *testing.T) {
	t.Parallel()
	type args struct {
		postalCode      string
		prefecture      Prefecture
		city            string
		streetAndNumber string
		buildingAndRoom string
	}

	tests := []struct {
		name string
		args args
		want *Address
	}{
		{
			name: "success",
			args: args{
				postalCode:      "1234567",
				prefecture:      PrefectureTokyo,
				city:            "city",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
			want: &Address{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "city",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := newAddress(tt.args.postalCode, tt.args.prefecture, tt.args.city, tt.args.streetAndNumber, tt.args.buildingAndRoom)
			if err != nil {
				t.Errorf("newAddress() error = %v, wantErr %v", err, nil)
				return
			}

			opt := cmp.AllowUnexported(Address{}, postalCode{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("newAddress() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestAddress_validate(t *testing.T) {
	t.Parallel()
	type fields struct {
		postalCode      postalCode
		prefecture      string
		city            string
		streetAndNumber string
		buildingAndRoom string
	}

	successTests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "city",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := &Address{
				postalCode:      tt.fields.postalCode,
				prefecture:      tt.fields.prefecture,
				city:            tt.fields.city,
				streetAndNumber: tt.fields.streetAndNumber,
				buildingAndRoom: tt.fields.buildingAndRoom,
			}

			if err := a.validate(); err != nil {
				t.Errorf("Address.validate() error = %v, wantErr %v", err, nil)
			}
		})
	}

	failureTests := []struct {
		name   string
		fields fields
	}{
		{
			name: "failure: city is empty",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
		},
		{
			name: "failure: street and number is empty",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "city",
				streetAndNumber: "",
				buildingAndRoom: "buildingAndRoom",
			},
		},
		{
			name: "failure: city is below the min length required",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
		},
		{
			name: "failure: street and number is below the min length required",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "city",
				streetAndNumber: "",
				buildingAndRoom: "buildingAndRoom",
			},
		},
		{
			name: "failure: city exceeds the max length allowed",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "citycitycitycitycitycitycitycitycitycitycitycitycity",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
		},
		{
			name: "failure: street and number exceeds the max length allowed",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "city",
				streetAndNumber: "streetAndNumberstreetAndNumberstreetAndNumberstreetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
		},
		{
			name: "failure: building and room exceeds the max length allowed",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "city",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoombuildingAndRoombuildingAndRoombuildingAndRoom",
			},
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := &Address{
				postalCode:      tt.fields.postalCode,
				prefecture:      tt.fields.prefecture,
				city:            tt.fields.city,
				streetAndNumber: tt.fields.streetAndNumber,
				buildingAndRoom: tt.fields.buildingAndRoom,
			}

			if err := a.validate(); err == nil {
				t.Errorf("Address.validate() error = %v, wantErr = is not nil", err)
			}
		})
	}
}

func TestAddress_PostalCode(t *testing.T) {
	t.Parallel()
	type fields struct {
		postalCode      postalCode
		prefecture      string
		city            string
		streetAndNumber string
		buildingAndRoom string
	}
	tests := []struct {
		name   string
		fields fields
		want   postalCode
	}{
		{
			name: "success",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "city",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
			want: postalCode{"1234567"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := Address{
				postalCode:      tt.fields.postalCode,
				prefecture:      tt.fields.prefecture,
				city:            tt.fields.city,
				streetAndNumber: tt.fields.streetAndNumber,
				buildingAndRoom: tt.fields.buildingAndRoom,
			}

			got := a.PostalCode()
			opt := cmp.AllowUnexported(postalCode{})

			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("Address.PostalCode() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestAddress_Prefecture(t *testing.T) {
	t.Parallel()
	type fields struct {
		postalCode      postalCode
		prefecture      string
		city            string
		streetAndNumber string
		buildingAndRoom string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "city",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
			want: prefectureTokyoValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := Address{
				postalCode:      tt.fields.postalCode,
				prefecture:      tt.fields.prefecture,
				city:            tt.fields.city,
				streetAndNumber: tt.fields.streetAndNumber,
				buildingAndRoom: tt.fields.buildingAndRoom,
			}

			got := a.Prefecture()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Address.Prefecture() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestAddress_City(t *testing.T) {
	t.Parallel()
	type fields struct {
		postalCode      postalCode
		prefecture      string
		city            string
		streetAndNumber string
		buildingAndRoom string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "city",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
			want: "city",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := Address{
				postalCode:      tt.fields.postalCode,
				prefecture:      tt.fields.prefecture,
				city:            tt.fields.city,
				streetAndNumber: tt.fields.streetAndNumber,
				buildingAndRoom: tt.fields.buildingAndRoom,
			}

			got := a.City()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Address.City() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestAddress_StreetAndNumber(t *testing.T) {
	t.Parallel()
	type fields struct {
		postalCode      postalCode
		prefecture      string
		city            string
		streetAndNumber string
		buildingAndRoom string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "city",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
			want: "streetAndNumber",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := Address{
				postalCode:      tt.fields.postalCode,
				prefecture:      tt.fields.prefecture,
				city:            tt.fields.city,
				streetAndNumber: tt.fields.streetAndNumber,
				buildingAndRoom: tt.fields.buildingAndRoom,
			}

			got := a.StreetAndNumber()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Address.StreetAndNumber() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestAddress_BuildingAndRoom(t *testing.T) {
	t.Parallel()
	type fields struct {
		postalCode      postalCode
		prefecture      string
		city            string
		streetAndNumber string
		buildingAndRoom string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "city",
				streetAndNumber: "streetAndNumber",
				buildingAndRoom: "buildingAndRoom",
			},
			want: "buildingAndRoom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := Address{
				postalCode:      tt.fields.postalCode,
				prefecture:      tt.fields.prefecture,
				city:            tt.fields.city,
				streetAndNumber: tt.fields.streetAndNumber,
				buildingAndRoom: tt.fields.buildingAndRoom,
			}

			got := a.BuildingAndRoom()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Address.BuildingAndRoom() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
