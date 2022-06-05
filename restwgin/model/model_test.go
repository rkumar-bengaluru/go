package model_test

import (
	"github.com/rkumar-bengaluru/go/restwgin/model"
	"testing"
)

func TestGetProduct(t *testing.T) {
	m := model.Product{ID: 1}
	m.GetProduct()
	if m.Name != "apple" {
		t.Errorf("Expected m.Name=apple whereas got m.Name=%v", m.Name)
	}

	if m.Price != 8.99 {
		t.Errorf("Expected m.Price=8.99 whereas got m.Price=%v", m.Price)
	}

}
