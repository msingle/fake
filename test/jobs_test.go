package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

func TestJobs(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		v := fake.Company()
		if v == "" {
			t.Errorf("Company failed with lang %s", lang)
		}

		v = fake.JobTitle()
		if v == "" {
			t.Errorf("JobTitle failed with lang %s", lang)
		}

		v = fake.Industry()
		if v == "" {
			t.Errorf("Industry failed with lang %s", lang)
		}
	}
}
