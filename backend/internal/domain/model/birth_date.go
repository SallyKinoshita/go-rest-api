package model

import (
	"time"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	ageMin = 18
)

type BirthDate struct {
	time.Time
}

func newBirthDate(birthDate, now time.Time) (BirthDate, error) {
	b := BirthDate{birthDate}

	if err := b.validate(now); err != nil {
		return BirthDate{}, err
	}

	return b, nil
}

func (b BirthDate) validate(now time.Time) error {
	age := b.CalAge(now)

	if err := validation.Validate(&age,
		validation.Required,
		validation.Min(ageMin),
	); err != nil {
		return cerror.Wrap(err, cerror.DetailUnprocessableEntityBirthDate)
	}

	return nil
}

func (b BirthDate) CalAge(now time.Time) int {
	age := now.Year() - b.Year()

	birthMonth := b.Month()
	birthDay := b.Day()
	currentMonth := now.Month()
	currentDay := now.Day()

	if b.IsLeapYear() {
		// 生年月日が閏年の場合
		if isYearLeap(now.Year()) {
			// 現在の年が閏年の場合、2月29日を誕生日として扱う
			if currentMonth < 2 || (currentMonth == 2 && currentDay < 29) {
				age--
			}
		} else {
			// 現在の年が閏年でない場合、2月28日を誕生日として扱う
			if currentMonth < 2 || (currentMonth == 2 && currentDay < 28) {
				age--
			}
		}
	} else {
		// 生年月日が閏年でない場合
		if currentMonth < birthMonth || (currentMonth == birthMonth && currentDay < birthDay) {
			age--
		}
	}

	return age
}

func (b BirthDate) IsLeapYear() bool {
	return b.Month() == 2 && b.Day() == 29
}

func isYearLeap(year int) bool {
	// 閏年の判定
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
