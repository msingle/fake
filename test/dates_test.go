package test

import (
	"testing"
	"time"

	"github.com/icrowley/fake"
)

func TestDates(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		fake.SetLang(lang)

		v := fake.WeekDay()
		if v == "" {
			t.Errorf("WeekDay failed with lang %s", lang)
		}

		v = fake.WeekDayShort()
		if v == "" {
			t.Errorf("WeekDayShort failed with lang %s", lang)
		}

		n := fake.WeekdayNum()
		if n < 0 || n > 7 {
			t.Errorf("WeekdayNum failed with lang %s", lang)
		}

		v = fake.Month()
		if v == "" {
			t.Errorf("Month failed with lang %s", lang)
		}

		v = fake.MonthShort()
		if v == "" {
			t.Errorf("MonthShort failed with lang %s", lang)
		}

		n = fake.MonthNum()
		if n < 0 || n > 31 {
			t.Errorf("MonthNum failed with lang %s", lang)
		}

		n = fake.Year(1950, 2020)
		if n < 1950 || n > 2020 {
			t.Errorf("Year failed with lang %s", lang)
		}
	}
}

func TestBirthdate(t *testing.T) {
	for age := 0; age <= 120; age += 1 {
		now := time.Now()
		birthdate := fake.Birthdate(age)
		diffDate := now.Sub(birthdate)
		calcAge := int(diffDate.Hours() / (365 * 24))
		diff := calcAge - age
		// TODO improve this to be more precise. Right now we don't need precision
		if diff > 1 {
			t.Errorf("Birthdate (%s) is not valid according to age (%d) = %d", birthdate, age, diff)
		}
	}
}
