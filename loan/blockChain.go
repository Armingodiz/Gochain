package loan

import "strings"

//RegisterLoan registers a Loan in our blockchain
func (b *Blockchain) RegisterBet(loan Loan) bool {
	loan.Name = strings.ToLower(loan.Name)
	loan.LoanID = strings.ToLower(loan.LoanID)
	b.PendingLoans = append(b.PendingLoans, loan)
	return true
}

//RegisterNode registers a node in our blockchain
func (b *Blockchain) RegisterNode(node string) bool {
	if !contains(b.NetworkNodes, node) {
		b.NetworkNodes = append(b.NetworkNodes, node)
	}
	return true
}

// function to check if new node is already added to nodes
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}


