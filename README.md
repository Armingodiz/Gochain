# Gochain
simple bank system for holding loans using blockchain

اونایی که جلوتر هستن internal  اند 

endpoints : 
		"Index",
		"GET",
		"/",
    
    
		"GetBlockchain",
		"GET",
		"/blockchain",



		"RegisterAndBroadcastNode",
		"POST", { the new node url => `json:"newnodeurl"` }
		"/register-and-broadcast-node",

  		  "RegisterNode",
	  	  "POST",
		    "/register-node",

		    "RegisterNodesBulk",
		    "POST",
		    "/register-nodes-bulk",



		"RegisterAndBroadcastLoan",
		"POST",
		"/loan/broadcast",
    
    		"RegisterLoan",
		    "POST",
		    "/loan",


		"Mine",
		"GET",
		"/mine",

  		"ReceiveNewBlock",
	  	"POST",
		  "/receive-new-block",


		"Consensus",
		"GET",
		"/consensus",


		"GetLoansForUser",
		"GET",
		"/user/{playerName}",
