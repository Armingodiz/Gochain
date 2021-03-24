package loan

type Loan struct {
	Name   string `json:"name"`
	LoanID string `json:"loanId"`
	Amount int    `json:"amount"`
}
