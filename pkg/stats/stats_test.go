package stats

import (
	"github.com/Payrav-1997/bank/v2/pkg/types"
		
	"testing"
	"reflect"
)


func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 1, Category: "food", Amount: 2_000_00},
		{ID: 1, Category: "auto", Amount: 4_000_00},
		{ID: 1, Category: "auto", Amount: 4_000_00},
		{ID: 1, Category: "fun", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 3_000_00,
		"food": 2_000_00,
		"fun":  5_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual %v", expected, result)
	}
}

func TestCategoriesAvg_nil(t *testing.T) {
	var payments []types.Payment
	result := CategoriesAvg(payments)

	if len(result) != 0 {
		t.Errorf("\n got > %v want > nil", result)
	}
}
func TestCategoriesAvg_one(t *testing.T) {
	payments := []types.Payment{
		{
			ID:       1,
			Category: "auto",
			Amount:   3_000_00,
		},
		{
			ID:       2,
			Category: "auto",
			Amount:   3_000_00,
		},
	}
	expected := map[types.Category]types.Money{
		"auto": 3_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\n got > %v want > nil", result)
	}
}