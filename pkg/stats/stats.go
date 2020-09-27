package stats

import (
	"github.com/Payrav-1997/bank/v2/pkg/types"
)

//CategoriesAvg считает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	category := map[types.Category]types.Money{}
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		category[payment.Category]++
	}

	for key := range categories {
		categories[key] /= category[key]
	}

	return categories
}
