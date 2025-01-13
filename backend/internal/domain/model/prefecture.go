package model

import (
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/bimap"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
)

type Prefecture int

const (
	PrefectureHokkaido Prefecture = iota + 1
	PrefectureAomori
	PrefectureIwate
	PrefectureMiyagi
	PrefectureAkita
	PrefectureYamagata
	PrefectureFukushima
	PrefectureIbaraki
	PrefectureTochigi
	PrefectureGunma
	PrefectureSaitama
	PrefectureChiba
	PrefectureTokyo
	PrefectureKanagawa
	PrefectureNiigata
	PrefectureToyama
	PrefectureIshikawa
	PrefectureFukui
	PrefectureYamanashi
	PrefectureNagano
	PrefectureGifu
	PrefectureShizuoka
	PrefectureAichi
	PrefectureMie
	PrefectureShiga
	PrefectureKyoto
	PrefectureOsaka
	PrefectureHyogo
	PrefectureNara
	PrefectureWakayama
	PrefectureTottori
	PrefectureShimane
	PrefectureOkayama
	PrefectureHiroshima
	PrefectureYamaguchi
	PrefectureTokushima
	PrefectureKagawa
	PrefectureEhime
	PrefectureKochi
	PrefectureFukuoka
	PrefectureSaga
	PrefectureNagasaki
	PrefectureKumamoto
	PrefectureOita
	PrefectureMiyazaki
	PrefectureKagoshima
	PrefectureOkinawa
)

const (
	prefectureHokkaidoValue  = "北海道"
	prefectureAomoriValue    = "青森県"
	prefectureIwateValue     = "岩手県"
	prefectureMiyagiValue    = "宮城県"
	prefectureAkitaValue     = "秋田県"
	prefectureYamagataValue  = "山形県"
	prefectureFukushimaValue = "福島県"
	prefectureIbarakiValue   = "茨城県"
	prefectureTochigiValue   = "栃木県"
	prefectureGunmaValue     = "群馬県"
	prefectureSaitamaValue   = "埼玉県"
	prefectureChibaValue     = "千葉県"
	prefectureTokyoValue     = "東京都"
	prefectureKanagawaValue  = "神奈川県"
	prefectureNiigataValue   = "新潟県"
	prefectureToyamaValue    = "富山県"
	prefectureIshikawaValue  = "石川県"
	prefectureFukuiValue     = "福井県"
	prefectureYamanashiValue = "山梨県"
	prefectureNaganoValue    = "長野県"
	prefectureGifuValue      = "岐阜県"
	prefectureShizuokaValue  = "静岡県"
	prefectureAichiValue     = "愛知県"
	prefectureMieValue       = "三重県"
	prefectureShigaValue     = "滋賀県"
	prefectureKyotoValue     = "京都府"
	prefectureOsakaValue     = "大阪府"
	prefectureHyogoValue     = "兵庫県"
	prefectureNaraValue      = "奈良県"
	prefectureWakayamaValue  = "和歌山県"
	prefectureTottoriValue   = "鳥取県"
	prefectureShimaneValue   = "島根県"
	prefectureOkayamaValue   = "岡山県"
	prefectureHiroshimaValue = "広島県"
	prefectureYamaguchiValue = "山口県"
	prefectureTokushimaValue = "徳島県"
	prefectureKagawaValue    = "香川県"
	prefectureEhimeValue     = "愛媛県"
	prefectureKochiValue     = "高知県"
	prefectureFukuokaValue   = "福岡県"
	prefectureSagaValue      = "佐賀県"
	prefectureNagasakiValue  = "長崎県"
	prefectureKumamotoValue  = "熊本県"
	prefectureOitaValue      = "大分県"
	prefectureMiyazakiValue  = "宮崎県"
	prefectureKagoshimaValue = "鹿児島県"
	prefectureOkinawaValue   = "沖縄県"
)

var prefectureBimap = bimap.NewMust(
	map[Prefecture]string{
		PrefectureHokkaido:  prefectureHokkaidoValue,
		PrefectureAomori:    prefectureAomoriValue,
		PrefectureIwate:     prefectureIwateValue,
		PrefectureMiyagi:    prefectureMiyagiValue,
		PrefectureAkita:     prefectureAkitaValue,
		PrefectureYamagata:  prefectureYamagataValue,
		PrefectureFukushima: prefectureFukushimaValue,
		PrefectureIbaraki:   prefectureIbarakiValue,
		PrefectureTochigi:   prefectureTochigiValue,
		PrefectureGunma:     prefectureGunmaValue,
		PrefectureSaitama:   prefectureSaitamaValue,
		PrefectureChiba:     prefectureChibaValue,
		PrefectureTokyo:     prefectureTokyoValue,
		PrefectureKanagawa:  prefectureKanagawaValue,
		PrefectureNiigata:   prefectureNiigataValue,
		PrefectureToyama:    prefectureToyamaValue,
		PrefectureIshikawa:  prefectureIshikawaValue,
		PrefectureFukui:     prefectureFukuiValue,
		PrefectureYamanashi: prefectureYamanashiValue,
		PrefectureNagano:    prefectureNaganoValue,
		PrefectureGifu:      prefectureGifuValue,
		PrefectureShizuoka:  prefectureShizuokaValue,
		PrefectureAichi:     prefectureAichiValue,
		PrefectureMie:       prefectureMieValue,
		PrefectureShiga:     prefectureShigaValue,
		PrefectureKyoto:     prefectureKyotoValue,
		PrefectureOsaka:     prefectureOsakaValue,
		PrefectureHyogo:     prefectureHyogoValue,
		PrefectureNara:      prefectureNaraValue,
		PrefectureWakayama:  prefectureWakayamaValue,
		PrefectureTottori:   prefectureTottoriValue,
		PrefectureShimane:   prefectureShimaneValue,
		PrefectureOkayama:   prefectureOkayamaValue,
		PrefectureHiroshima: prefectureHiroshimaValue,
		PrefectureYamaguchi: prefectureYamaguchiValue,
		PrefectureTokushima: prefectureTokushimaValue,
		PrefectureKagawa:    prefectureKagawaValue,
		PrefectureEhime:     prefectureEhimeValue,
		PrefectureKochi:     prefectureKochiValue,
		PrefectureFukuoka:   prefectureFukuokaValue,
		PrefectureSaga:      prefectureSagaValue,
		PrefectureNagasaki:  prefectureNagasakiValue,
		PrefectureKumamoto:  prefectureKumamotoValue,
		PrefectureOita:      prefectureOitaValue,
		PrefectureMiyazaki:  prefectureMiyazakiValue,
		PrefectureKagoshima: prefectureKagoshimaValue,
		PrefectureOkinawa:   prefectureOkinawaValue,
	},
)

func (p Prefecture) Value() (string, error) {
	value, ok := prefectureBimap.GetForwardMap()[p]
	if !ok {
		return "", cerror.New("failed to prefecture value", cerror.DetailInternalServerError)
	}

	return value, nil
}

func PrefectureKeyByValue(value string) (Prefecture, error) {
	key, ok := prefectureBimap.GetReverseMap()[value]
	if !ok {
		return 0, cerror.New("failed to prefecture key", cerror.DetailInternalServerError)
	}

	return key, nil
}
