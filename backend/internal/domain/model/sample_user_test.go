package model

import (
	"testing"
	"time"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cuuid"
	"github.com/google/go-cmp/cmp"
)

func TestNewSampleUser(t *testing.T) {
	t.Parallel()
	type args struct {
		sampleUserID    cuuid.CUUID
		firstName       string
		lastName        string
		firstNameKana   string
		lastNameKana    string
		emailAddress    string
		tel             string
		birthDate       time.Time
		postalCode      string
		prefecture      Prefecture
		city            string
		streetAndNumber string
		buildingAndRoom string
		now             time.Time
	}
	tests := []struct {
		name string
		args args
		want *SampleUser
	}{
		{
			name: "success",
			args: args{
				sampleUserID:    cuuid.TestCUUID(),
				firstName:       "太郎",
				lastName:        "山田",
				firstNameKana:   "タロウ",
				lastNameKana:    "ヤマダ",
				emailAddress:    "sample@exsample.com",
				tel:             "+11234567890",
				birthDate:       birthDate,
				postalCode:      "1234567",
				prefecture:      PrefectureTokyo,
				city:            "渋谷区",
				streetAndNumber: "道玄坂1-2-3",
				buildingAndRoom: "ビル101",
				now:             now,
			},
			want: &SampleUser{
				id: cuuid.TestCUUID(),
				name: Name{
					first:     "太郎",
					last:      "山田",
					firstKana: "タロウ",
					lastKana:  "ヤマダ",
				},
				emailAddress: EmailAddress{"sample@exsample.com"},
				tel:          Tel{"+11234567890"},
				birthDate:    BirthDate{birthDate},
				address: Address{
					postalCode:      postalCode{"1234567"},
					prefecture:      prefectureTokyoValue,
					city:            "渋谷区",
					streetAndNumber: "道玄坂1-2-3",
					buildingAndRoom: "ビル101",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewSampleUser(tt.args.sampleUserID, tt.args.firstName, tt.args.lastName, tt.args.firstNameKana, tt.args.lastNameKana, tt.args.emailAddress, tt.args.tel, tt.args.birthDate, tt.args.postalCode, tt.args.prefecture, tt.args.city, tt.args.streetAndNumber, tt.args.buildingAndRoom, tt.args.now)
			if err != nil {
				t.Errorf("NewSampleUser() error = %v, wantErr %v", err, nil)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewSampleUser() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestSampleUser_Update(t *testing.T) {
	t.Parallel()
	type fields struct {
		id           cuuid.CUUID
		name         Name
		emailAddress EmailAddress
		tel          Tel
		birthDate    BirthDate
		address      Address
	}
	type args struct {
		firstName       string
		lastName        string
		firstNameKana   string
		lastNameKana    string
		emailAddress    string
		postalCode      string
		prefecture      Prefecture
		city            string
		streetAndNumber string
		buildingAndRoom string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				id: cuuid.TestCUUID(),
				name: Name{
					first:     "太郎",
					last:      "山田",
					firstKana: "タロウ",
					lastKana:  "ヤマダ",
				},
				emailAddress: EmailAddress{"sample@exsample.com"},
				tel:          Tel{"+11234567890"},
				birthDate:    BirthDate{birthDate},
				address: Address{
					postalCode:      postalCode{"1234567"},
					prefecture:      prefectureTokyoValue,
					city:            "渋谷区",
					streetAndNumber: "道玄坂1-2-3",
					buildingAndRoom: "ビル101",
				},
			},
			args: args{
				firstName:       "小太郎",
				lastName:        "山田",
				firstNameKana:   "コタロウ",
				lastNameKana:    "ヤマダ",
				emailAddress:    "sample2@exsample.com",
				postalCode:      "6028570",
				prefecture:      PrefectureKyoto,
				city:            "京都市",
				streetAndNumber: "上京区下立売通新町西入薮ノ内町",
				buildingAndRoom: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &SampleUser{
				id:           tt.fields.id,
				name:         tt.fields.name,
				emailAddress: tt.fields.emailAddress,
				tel:          tt.fields.tel,
				birthDate:    tt.fields.birthDate,
				address:      tt.fields.address,
			}
			if err := s.Update(tt.args.firstName, tt.args.lastName, tt.args.firstNameKana, tt.args.lastNameKana, tt.args.emailAddress, tt.args.postalCode, tt.args.prefecture, tt.args.city, tt.args.streetAndNumber, tt.args.buildingAndRoom); err != nil {
				t.Errorf("Update() error = %v, wantErr %v", err, nil)
			}
		})
	}
}

func TestSampleUser_ID(t *testing.T) {
	t.Parallel()
	type fields struct {
		id cuuid.CUUID
	}

	tests := []struct {
		name   string
		fields fields
		want   cuuid.CUUID
	}{
		{
			name: "success",
			fields: fields{
				id: cuuid.TestCUUID(),
			},
			want: cuuid.TestCUUID(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := SampleUser{
				id: tt.fields.id,
			}
			got := s.ID()
			opt := cmp.AllowUnexported(cuuid.CUUID{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("SampleUser.ID() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestSampleUser_Name(t *testing.T) {
	t.Parallel()
	type fields struct {
		name Name
	}

	tests := []struct {
		name   string
		fields fields
		want   Name
	}{
		{
			name: "success",
			fields: fields{
				name: Name{
					first:     "太郎",
					last:      "山田",
					firstKana: "タロウ",
					lastKana:  "ヤマダ",
				},
			},
			want: Name{
				first:     "太郎",
				last:      "山田",
				firstKana: "タロウ",
				lastKana:  "ヤマダ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := SampleUser{
				name: tt.fields.name,
			}
			got := s.Name()
			opt := cmp.AllowUnexported(Name{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("SampleUser.Name() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestSampleUser_EmailAddress(t *testing.T) {
	t.Parallel()
	type fields struct {
		emailAddress EmailAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   EmailAddress
	}{
		{
			name: "success",
			fields: fields{
				emailAddress: EmailAddress{"sample@exsample.com"},
			},
			want: EmailAddress{"sample@exsample.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := SampleUser{
				emailAddress: tt.fields.emailAddress,
			}
			got := s.EmailAddress()
			opt := cmp.AllowUnexported(EmailAddress{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("SampleUser.EmailAddress() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestSampleUser_Tel(t *testing.T) {
	t.Parallel()
	type fields struct {
		tel Tel
	}
	tests := []struct {
		name   string
		fields fields
		want   Tel
	}{
		{
			name: "success",
			fields: fields{
				tel: Tel{"+11234567890"},
			},
			want: Tel{"+11234567890"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := SampleUser{
				tel: tt.fields.tel,
			}
			got := s.Tel()
			opt := cmp.AllowUnexported(Tel{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("SampleUser.Tel() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestSampleUser_BirthDate(t *testing.T) {
	t.Parallel()
	type fields struct {
		birthDate BirthDate
	}
	tests := []struct {
		name   string
		fields fields
		want   BirthDate
	}{
		{
			name: "success",
			fields: fields{
				birthDate: BirthDate{birthDate},
			},
			want: BirthDate{birthDate},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := SampleUser{
				birthDate: tt.fields.birthDate,
			}
			got := s.BirthDate()
			opt := cmp.AllowUnexported(BirthDate{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("SampleUser.BirthDate() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestSampleUser_Address(t *testing.T) {
	t.Parallel()
	type fields struct {
		address Address
	}
	tests := []struct {
		name   string
		fields fields
		want   Address
	}{
		{
			name: "success",
			fields: fields{
				address: Address{
					postalCode:      postalCode{"1234567"},
					prefecture:      prefectureTokyoValue,
					city:            "渋谷区",
					streetAndNumber: "道玄坂1-2-3",
					buildingAndRoom: "ビル101",
				},
			},
			want: Address{
				postalCode:      postalCode{"1234567"},
				prefecture:      prefectureTokyoValue,
				city:            "渋谷区",
				streetAndNumber: "道玄坂1-2-3",
				buildingAndRoom: "ビル101",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := SampleUser{
				address: tt.fields.address,
			}
			got := s.Address()
			opt := cmp.AllowUnexported(Address{}, postalCode{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("SampleUser.Address() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
