package loan

import (
	"strings"
	"time"
)

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

//CreateNewBlock create and add new block to our blockchain (this method will be called after validation)
func (b *Blockchain) CreateNewBlock(nonce int, previousBlockHash string, hash string) Block {
	newBlock := Block{
		Index:     len(b.Chain) + 1,
		Loans:     b.PendingLoans,
		Timestamp: time.Now().UnixNano(),
		Nonce:     nonce,
		Hash:      hash, PreviousBlockHash: previousBlockHash}
	b.PendingLoans = Loans{}
	b.Chain = append(b.Chain, newBlock)
	return newBlock
}

//GetLastBlock
func (b *Blockchain) GetLastBlock() Block {
	return b.Chain[len(b.Chain)-1]
}
