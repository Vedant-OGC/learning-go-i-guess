package expense_tracker

type Category int

const (
	Food Category = iota
	Travel
	Utilities
	Entertainment
)

type Expense struct {
	ID       int      `json:"id"`
	Amount   float64  `json:"amount"`
	Category Category `json:"category"`
	Note     string   `json:"note"`
}
