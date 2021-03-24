package loan

import (
	"github.com/ArminGodiz/Gochain/logger"
	"github.com/gorilla/mux"
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

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"GetBlockchain",
		"GET",
		"/blockchain",
		controller.GetBlockchain,
	},
	Route{
		"RegisterAndBroadcastNode",
		"POST",
		"/register-and-broadcast-node",
		controller.RegisterAndBroadcastNode,
	},
	Route{
		"RegisterNode",
		"POST",
		"/register-node",
		controller.RegisterNode,
	},
	Route{
		"RegisterNodesBulk",
		"POST",
		"/register-nodes-bulk",
		controller.RegisterNodesBulk,
	},
	/*
		will pass all the current nodes of the network to the new node, so they it will also have them in its internal list of nodes
	*/
	Route{
		"RegisterLoan",
		"POST",
		"/loan",
		controller.RegisterLoan,
	},
	Route{
		"RegisterAndBroadcastLoan",
		"POST",
		"/loan/broadcast",
		controller.RegisterAndBroadcastLoan,
	},
	Route{
		"Mine",
		"GET",
		"/mine",
		controller.Mine,
	},
	Route{
		"ReceiveNewBlock",
		"POST",
		"/receive-new-block",
		controller.ReceiveNewBlock,
	},
	Route{
		"Consensus",
		"GET",
		"/consensus",
		controller.Consensus,
	},
	Route{
		"GetLoansForUser",
		"GET",
		"/user/{playerName}",
		controller.GetLoansForUser,
	},
}

//NewRouter configures a new router to the API
func NewRouter(nodeAddress string) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	controller.currentNodeURL = "http://localhost:" + nodeAddress

	// create Genesis block
	controller.blockchain.CreateNewBlock(100, "0", "0")

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}