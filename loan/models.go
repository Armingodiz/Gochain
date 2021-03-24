package loan

// Loan :
type Loan struct {
	Name   string `json:"name"`
	LoanID string `json:"loanId"`
	Amount int    `json:"amount"`
}

//Loans is an array of Loan
type Loans []Loan

//Block :
type Block struct {
	Index             int    `json:"index"`
	Timestamp         int64  `json:"timestamp"`
	Loans             Loans  `json:"loans"`
	Nonce             int    `json:"nonce"`
	Hash              string `json:"hash"`
	PreviousBlockHash string `json:"previousBlockHash"`
}

