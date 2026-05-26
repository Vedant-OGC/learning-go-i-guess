package expense_tracker

import "errors"

type Category int

const (
	Food Category = iota
	Travel
	Utilities
	Entertainment
)

func (c Category) String() string {
	switch c {
	case Food:
		return "Food"
	case Travel:
		return "Travel"
	case Utilities:
		return "Utilities"
	case Entertainment:
		return "Entertainment"
	default:
		return "Other"
	}
}

type Expense struct {
	ID       int      `json:"id"`
	Amount   float64  `json:"amount"`
	Category Category `json:"category"`
	Note     string   `json:"note"`
}

type ExpenseBook struct {
	Expenses []Expense
}

func (eb *ExpenseBook) Add(amount float64, cat Category, note string) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	id := len(eb.Expenses) + 1
	eb.Expenses = append(eb.Expenses, Expense{ID: id, Amount: amount, Category: cat, Note: note})
	return nil
}

func (eb *ExpenseBook) Delete(id int) bool {
	for i, e := range eb.Expenses {
		if e.ID == id {
			eb.Expenses = append(eb.Expenses[:i], eb.Expenses[i+1:]...)
			return true
		}
	}
	return false
}
