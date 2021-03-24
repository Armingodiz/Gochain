package loan

import (
	"crypto/sha256"
	"encoding/base64"
	"strconv"
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

//HashBlock Create hash id for block base on previous block , data and nonce of block
func (b *Blockchain) HashBlock(previousBlockHash string, currentBlockData string, nonce int) string {
	h := sha256.New()
	strToHash := previousBlockHash + currentBlockData + strconv.Itoa(nonce)
	h.Write([]byte(strToHash))
	hashed := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return hashed
}

//CheckNewBlockHash checks block validation by checking previous block hash and its index in chain :

func (b *Blockchain) CheckNewBlockHash(newBlock Block) bool {
	lastBlock := b.GetLastBlock()
	correctHash := lastBlock.Hash == newBlock.PreviousBlockHash
	correctIndex := (lastBlock.Index + 1) == newBlock.Index

	return (correctHash && correctIndex)
}

//ProofOfWork consists of finding a nonce (an integer in our case) that, combined with all the other data in the block, will return a hash code which begins with “0000”.
func (b *Blockchain) ProofOfWork(previousBlockHash string, currentBlockData string) int {
	nonce := -1
	inputFmt := ""
	for inputFmt != "0000" {
		nonce = nonce + 1
		hash := b.HashBlock(previousBlockHash, currentBlockData, nonce)
		inputFmt = hash[0:4]
	}
	return nonce
}
