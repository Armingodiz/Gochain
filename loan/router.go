package loan

import (
	"net/http"
)

var controller = &Controller{
	blockchain: &Blockchain{
		Chain:        Blocks{},
		PendingLoans: Loans{},
		NetworkNodes: []string{}}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
