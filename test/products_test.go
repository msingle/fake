package test

import (
	"testing"

	"github.com/syscrusher/fake"
)

func TestProducts(t *testing.T) {
	for _, lang := range fake.GetLangs() {
		err := fake.SetLang(lang)
		if err != nil {
			t.Errorf("Could not set language %s", lang)
		}

		v := fake.Brand()
		if v == "" {
			t.Errorf("Brand failed with lang %s", lang)
		}

		v = fake.ProductName()
		if v == "" {
			t.Errorf("ProductName failed with lang %s", lang)
		}

		v = fake.Product()
		if v == "" {
			t.Errorf("Product failed with lang %s", lang)
		}

		v = fake.Model()
		if v == "" {
			t.Errorf("Model failed with lang %s", lang)
		}
	}
}
