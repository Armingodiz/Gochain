package loan

import "strings"

//RegisterLoan registers a Loan in our blockchain
func (b *Blockchain) RegisterBet(loan Loan) bool {
	loan.Name = strings.ToLower(loan.Name)
	loan.LoanID = strings.ToLower(loan.LoanID)
	b.PendingLoans = append(b.PendingLoans, loan)
	return true
}


