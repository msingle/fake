package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

func TestPersonal(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		v := fake.Gender()
		if v == "" {
			t.Errorf("Gender failed with lang %s", lang)
		}

		v = fake.GenderAbbrev()
		if v == "" {
			t.Errorf("GenderAbbrev failed with lang %s", lang)
		}

		v = fake.Language()
		if v == "" {
			t.Errorf("Language failed with lang %s", lang)
		}
	}
}
