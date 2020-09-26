package stats

import (
	"fmt"

	"github.com/Payrav-1997/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   53_00,
			Category: "Cafe",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   51_00,
			Category: "Cafe",
			Status:   types.StatusOk,
		},
		{
			ID:       1,
			Amount:   52_00,
			Category: "Cafe",
			Status:   types.StatusFail,
		},
	}

	fmt.Println(Avg(payments))
	//Output: 5200
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "cafe",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   20_000_00,
			Category: "Home",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "restaurant",
			Status:   types.StatusFail,
		},
	}

	result := types.Category("cafe")
	totalInCategory := TotalInCategory(payments, result)
	fmt.Println(totalInCategory)
	//Output:  1000000

}
