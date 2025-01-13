package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrefecture_Value(t *testing.T) {
	t.Parallel()

	successTests := []struct {
		name string
		p    Prefecture
		want string
	}{
		{
			name: "success: prefecture is Hokkaido",
			p:    PrefectureHokkaido,
			want: prefectureHokkaidoValue,
		},
		{
			name: "success: prefecture is Aomori",
			p:    PrefectureAomori,
			want: prefectureAomoriValue,
		},
		{
			name: "success: prefecture is Iwate",
			p:    PrefectureIwate,
			want: prefectureIwateValue,
		},
		{
			name: "success: prefecture is Miyagi",
			p:    PrefectureMiyagi,
			want: prefectureMiyagiValue,
		},
		{
			name: "success: prefecture is Akita",
			p:    PrefectureAkita,
			want: prefectureAkitaValue,
		},
		{
			name: "success: prefecture is Yamagata",
			p:    PrefectureYamagata,
			want: prefectureYamagataValue,
		},
		{
			name: "success: prefecture is Fukushima",
			p:    PrefectureFukushima,
			want: prefectureFukushimaValue,
		},
		{
			name: "success: prefecture is Ibaraki",
			p:    PrefectureIbaraki,
			want: prefectureIbarakiValue,
		},
		{
			name: "success: prefecture is Tochigi",
			p:    PrefectureTochigi,
			want: prefectureTochigiValue,
		},
		{
			name: "success: prefecture is Gunma",
			p:    PrefectureGunma,
			want: prefectureGunmaValue,
		},
		{
			name: "success: prefecture is Saitama",
			p:    PrefectureSaitama,
			want: prefectureSaitamaValue,
		},
		{
			name: "success: prefecture is Chiba",
			p:    PrefectureChiba,
			want: prefectureChibaValue,
		},
		{
			name: "success: prefecture is Tokyo",
			p:    PrefectureTokyo,
			want: prefectureTokyoValue,
		},
		{
			name: "success: prefecture is Kanagawa",
			p:    PrefectureKanagawa,
			want: prefectureKanagawaValue,
		},
		{
			name: "success: prefecture is Niigata",
			p:    PrefectureNiigata,
			want: prefectureNiigataValue,
		},
		{
			name: "success: prefecture is Toyama",
			p:    PrefectureToyama,
			want: prefectureToyamaValue,
		},
		{
			name: "success: prefecture is Ishikawa",
			p:    PrefectureIshikawa,
			want: prefectureIshikawaValue,
		},
		{
			name: "success: prefecture is Fukui",
			p:    PrefectureFukui,
			want: prefectureFukuiValue,
		},
		{
			name: "success: prefecture is Yamanashi",
			p:    PrefectureYamanashi,
			want: prefectureYamanashiValue,
		},
		{
			name: "success: prefecture is Nagano",
			p:    PrefectureNagano,
			want: prefectureNaganoValue,
		},
		{
			name: "success: prefecture is Gifu",
			p:    PrefectureGifu,
			want: prefectureGifuValue,
		},
		{
			name: "success: prefecture is Shizuoka",
			p:    PrefectureShizuoka,
			want: prefectureShizuokaValue,
		},
		{
			name: "success: prefecture is Aichi",
			p:    PrefectureAichi,
			want: prefectureAichiValue,
		},
		{
			name: "success: prefecture is Mie",
			p:    PrefectureMie,
			want: prefectureMieValue,
		},
		{
			name: "success: prefecture is Shiga",
			p:    PrefectureShiga,
			want: prefectureShigaValue,
		},
		{
			name: "success: prefecture is Kyoto",
			p:    PrefectureKyoto,
			want: prefectureKyotoValue,
		},
		{
			name: "success: prefecture is Osaka",
			p:    PrefectureOsaka,
			want: prefectureOsakaValue,
		},
		{
			name: "success: prefecture is Hyogo",
			p:    PrefectureHyogo,
			want: prefectureHyogoValue,
		},
		{
			name: "success: prefecture is Nara",
			p:    PrefectureNara,
			want: prefectureNaraValue,
		},
		{
			name: "success: prefecture is Wakayama",
			p:    PrefectureWakayama,
			want: prefectureWakayamaValue,
		},
		{
			name: "success: prefecture is Tottori",
			p:    PrefectureTottori,
			want: prefectureTottoriValue,
		},
		{
			name: "success: prefecture is Shimane",
			p:    PrefectureShimane,
			want: prefectureShimaneValue,
		},
		{
			name: "success: prefecture is Okayama",
			p:    PrefectureOkayama,
			want: prefectureOkayamaValue,
		},
		{
			name: "success: prefecture is Hiroshima",
			p:    PrefectureHiroshima,
			want: prefectureHiroshimaValue,
		},
		{
			name: "success: prefecture is Yamaguchi",
			p:    PrefectureYamaguchi,
			want: prefectureYamaguchiValue,
		},
		{
			name: "success: prefecture is Tokushima",
			p:    PrefectureTokushima,
			want: prefectureTokushimaValue,
		},
		{
			name: "success: prefecture is Kagawa",
			p:    PrefectureKagawa,
			want: prefectureKagawaValue,
		},
		{
			name: "success: prefecture is Ehime",
			p:    PrefectureEhime,
			want: prefectureEhimeValue,
		},
		{
			name: "success: prefecture is Kochi",
			p:    PrefectureKochi,
			want: prefectureKochiValue,
		},
		{
			name: "success: prefecture is Fukuoka",
			p:    PrefectureFukuoka,
			want: prefectureFukuokaValue,
		},
		{
			name: "success: prefecture is Saga",
			p:    PrefectureSaga,
			want: prefectureSagaValue,
		},
		{
			name: "success: prefecture is Nagasaki",
			p:    PrefectureNagasaki,
			want: prefectureNagasakiValue,
		},
		{
			name: "success: prefecture is Kumamoto",
			p:    PrefectureKumamoto,
			want: prefectureKumamotoValue,
		},
		{
			name: "success: prefecture is Oita",
			p:    PrefectureOita,
			want: prefectureOitaValue,
		},
		{
			name: "success: prefecture is Miyazaki",
			p:    PrefectureMiyazaki,
			want: prefectureMiyazakiValue,
		},
		{
			name: "success: prefecture is Kagoshima",
			p:    PrefectureKagoshima,
			want: prefectureKagoshimaValue,
		},
		{
			name: "success: prefecture is Okinawa",
			p:    PrefectureOkinawa,
			want: prefectureOkinawaValue,
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.p.Value()
			if err != nil {
				t.Errorf("Prefecture.Value() error = %v, wantErr %v", err, nil)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Prefecture.Value() mismatch (-got +want):\n%s", diff)
			}
		})
	}

	failureTests := []struct {
		name string
		p    Prefecture
		want string
	}{
		{
			name: "failure: prefecture is not defined",
			p:    0,
			want: "",
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.p.Value()
			if err == nil {
				t.Errorf("Prefecture.Value() error = %v, wantErr %v", nil, err)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Prefecture.Value() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func TestPrefectureKeyByValue(t *testing.T) {
	t.Parallel()
	type args struct {
		value string
	}

	successTests := []struct {
		name string
		args args
		want Prefecture
	}{
		{
			name: "success: prefecture is Hokkaido",
			args: args{value: prefectureHokkaidoValue},
			want: PrefectureHokkaido,
		},
		{
			name: "success: prefecture is Aomori",
			args: args{value: prefectureAomoriValue},
			want: PrefectureAomori,
		},
		{
			name: "success: prefecture is Iwate",
			args: args{value: prefectureIwateValue},
			want: PrefectureIwate,
		},
		{
			name: "success: prefecture is Miyagi",
			args: args{value: prefectureMiyagiValue},
			want: PrefectureMiyagi,
		},
		{
			name: "success: prefecture is Akita",
			args: args{value: prefectureAkitaValue},
			want: PrefectureAkita,
		},
		{
			name: "success: prefecture is Yamagata",
			args: args{value: prefectureYamagataValue},
			want: PrefectureYamagata,
		},
		{
			name: "success: prefecture is Fukushima",
			args: args{value: prefectureFukushimaValue},
			want: PrefectureFukushima,
		},
		{
			name: "success: prefecture is Ibaraki",
			args: args{value: prefectureIbarakiValue},
			want: PrefectureIbaraki,
		},
		{
			name: "success: prefecture is Tochigi",
			args: args{value: prefectureTochigiValue},
			want: PrefectureTochigi,
		},
		{
			name: "success: prefecture is Gunma",
			args: args{value: prefectureGunmaValue},
			want: PrefectureGunma,
		},
		{
			name: "success: prefecture is Saitama",
			args: args{value: prefectureSaitamaValue},
			want: PrefectureSaitama,
		},
		{
			name: "success: prefecture is Chiba",
			args: args{value: prefectureChibaValue},
			want: PrefectureChiba,
		},
		{
			name: "success: prefecture is Tokyo",
			args: args{value: prefectureTokyoValue},
			want: PrefectureTokyo,
		},
		{
			name: "success: prefecture is Kanagawa",
			args: args{value: prefectureKanagawaValue},
			want: PrefectureKanagawa,
		},
		{
			name: "success: prefecture is Niigata",
			args: args{value: prefectureNiigataValue},
			want: PrefectureNiigata,
		},
		{
			name: "success: prefecture is Toyama",
			args: args{value: prefectureToyamaValue},
			want: PrefectureToyama,
		},
		{
			name: "success: prefecture is Ishikawa",
			args: args{value: prefectureIshikawaValue},
			want: PrefectureIshikawa,
		},
		{
			name: "success: prefecture is Fukui",
			args: args{value: prefectureFukuiValue},
			want: PrefectureFukui,
		},
		{
			name: "success: prefecture is Yamanashi",
			args: args{value: prefectureYamanashiValue},
			want: PrefectureYamanashi,
		},
		{
			name: "success: prefecture is Nagano",
			args: args{value: prefectureNaganoValue},
			want: PrefectureNagano,
		},
		{
			name: "success: prefecture is Gifu",
			args: args{value: prefectureGifuValue},
			want: PrefectureGifu,
		},
		{
			name: "success: prefecture is Shizuoka",
			args: args{value: prefectureShizuokaValue},
			want: PrefectureShizuoka,
		},
		{
			name: "success: prefecture is Aichi",
			args: args{value: prefectureAichiValue},
			want: PrefectureAichi,
		},
		{
			name: "success: prefecture is Mie",
			args: args{value: prefectureMieValue},
			want: PrefectureMie,
		},
		{
			name: "success: prefecture is Shiga",
			args: args{value: prefectureShigaValue},
			want: PrefectureShiga,
		},
		{
			name: "success: prefecture is Kyoto",
			args: args{value: prefectureKyotoValue},
			want: PrefectureKyoto,
		},
		{
			name: "success: prefecture is Osaka",
			args: args{value: prefectureOsakaValue},
			want: PrefectureOsaka,
		},
		{
			name: "success: prefecture is Hyogo",
			args: args{value: prefectureHyogoValue},
			want: PrefectureHyogo,
		},
		{
			name: "success: prefecture is Nara",
			args: args{value: prefectureNaraValue},
			want: PrefectureNara,
		},
		{
			name: "success: prefecture is Wakayama",
			args: args{value: prefectureWakayamaValue},
			want: PrefectureWakayama,
		},
		{
			name: "success: prefecture is Tottori",
			args: args{value: prefectureTottoriValue},
			want: PrefectureTottori,
		},
		{
			name: "success: prefecture is Shimane",
			args: args{value: prefectureShimaneValue},
			want: PrefectureShimane,
		},
		{
			name: "success: prefecture is Okayama",
			args: args{value: prefectureOkayamaValue},
			want: PrefectureOkayama,
		},
		{
			name: "success: prefecture is Hiroshima",
			args: args{value: prefectureHiroshimaValue},
			want: PrefectureHiroshima,
		},
		{
			name: "success: prefecture is Yamaguchi",
			args: args{value: prefectureYamaguchiValue},
			want: PrefectureYamaguchi,
		},
		{
			name: "success: prefecture is Tokushima",
			args: args{value: prefectureTokushimaValue},
			want: PrefectureTokushima,
		},
		{
			name: "success: prefecture is Kagawa",
			args: args{value: prefectureKagawaValue},
			want: PrefectureKagawa,
		},
		{
			name: "success: prefecture is Ehime",
			args: args{value: prefectureEhimeValue},
			want: PrefectureEhime,
		},
		{
			name: "success: prefecture is Kochi",
			args: args{value: prefectureKochiValue},
			want: PrefectureKochi,
		},
		{
			name: "success: prefecture is Fukuoka",
			args: args{value: prefectureFukuokaValue},
			want: PrefectureFukuoka,
		},
		{
			name: "success: prefecture is Saga",
			args: args{value: prefectureSagaValue},
			want: PrefectureSaga,
		},
		{
			name: "success: prefecture is Nagasaki",
			args: args{value: prefectureNagasakiValue},
			want: PrefectureNagasaki,
		},
		{
			name: "success: prefecture is Kumamoto",
			args: args{value: prefectureKumamotoValue},
			want: PrefectureKumamoto,
		},
		{
			name: "success: prefecture is Oita",
			args: args{value: prefectureOitaValue},
			want: PrefectureOita,
		},
		{
			name: "success: prefecture is Miyazaki",
			args: args{value: prefectureMiyazakiValue},
			want: PrefectureMiyazaki,
		},
		{
			name: "success: prefecture is Kagoshima",
			args: args{value: prefectureKagoshimaValue},
			want: PrefectureKagoshima,
		},
		{
			name: "success: prefecture is Okinawa",
			args: args{value: prefectureOkinawaValue},
			want: PrefectureOkinawa,
		},
	}
	for _, tt := range successTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := PrefectureKeyByValue(tt.args.value)
			if err != nil {
				t.Errorf("PrefectureKeyByValue() error = %v, wantErr %v", err, nil)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("PrefectureKeyByValue() mismatch (-got +want):\n%s", diff)
			}
		})
	}

	failureTests := []struct {
		name string
		args args
		want Prefecture
	}{
		{
			name: "failure: prefecture is empty",
			args: args{value: ""},
			want: 0,
		},
		{
			name: "failure: prefecture is not found",
			args: args{value: "unknown"},
			want: 0,
		},
	}
	for _, tt := range failureTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := PrefectureKeyByValue(tt.args.value)
			if err == nil {
				t.Errorf("PrefectureKeyByValue() error = %v, wantErr %v", nil, err)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("PrefectureKeyByValue() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
