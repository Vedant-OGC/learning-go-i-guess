package expense_tracker

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
