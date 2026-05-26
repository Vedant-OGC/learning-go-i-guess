package expense_tracker

func CalculateTotal(expenses []Expense) float64 {
	total := 0.0
	for _, e := range expenses {
		total += e.Amount
	}
	return total
}

func FilterByCategory(expenses []Expense, cat Category) []Expense {
	var filtered []Expense
	for _, e := range expenses {
		if e.Category == cat {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
