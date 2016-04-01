package test

import (
	"testing"
	"time"

	"github.com/syscrusher/fake"
)

func TestDates(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		from, _ := time.Parse("2006-01-02T15:04:05", "2016-01-01T00:00:00")
		to, _ := time.Parse("2006-01-02T15:04:05", "2016-01-31T23:59:59")
		d := fake.Time(from, to)
		if !from.Equal(d) && !from.Before(d) {
			t.Errorf("Time before from date %s", lang)
		}
		if !to.Equal(d) && !to.After(d) {
			t.Errorf("Time after %s to date %s", d, lang)
		}

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
	for age := 0; age <= 120; age++ {
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
