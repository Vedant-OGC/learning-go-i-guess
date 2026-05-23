package expense_tracker

func CalculateTotal(expenses []Expense) float64 {
	total := 0.0
	for _, e := range expenses {
		total += e.Amount
	}
	return total
}
