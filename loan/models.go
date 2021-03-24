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
	Nonce             int    `json:"nonce"` // will be used for mining
	Hash              string `json:"hash"`
	PreviousBlockHash string `json:"previousBlockHash"`
}

//Blocks is an array of Block
type Blocks []Block

//Blockchain :
type Blockchain struct {
	Chain        Blocks   `json:"chain"`
	PendingLoans Loans    `json:"pending_loans"`
	NetworkNodes []string `json:"network_nodes"`
}

//BlockData is used in hash calculations
type BlockData struct {
	Index string
	Loans Loans
}
