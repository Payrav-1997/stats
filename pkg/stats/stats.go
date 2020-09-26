package stats

import (
	"github.com/Payrav-1997/bank/v2/pkg/types"
)

//Avg рассчитает среднюю сумму
func Avg(payments []types.Payment) types.Money {

	var sum = types.Money(0)
	var i = 0
	for _, card := range payments {
		if card.Status != types.StatusFail {
			sum += card.Amount
			i++
		}
	}
	return sum / types.Money(i)
}

//TotalInCategory находит сумму в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var t = types.Money(0)

	for _, card := range payments {
		if card.Status != types.StatusFail {
			if category == card.Category {
				t += card.Amount
			}
		}
	}

	return t
}
