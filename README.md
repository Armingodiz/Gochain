# Gochain
simple bank system for holding loans using blockchain 

# Installation
After cloning repo you can run `go run main.go "port"` . Your node will be listenning on localhost:port (each node can be considered as a Bank branch).
If you want your other nodes to be connected to the same network you must use RegisterAndBroadcastNode endpoint to send one of current nodes in network in your post request.

## Dependencies
name     | repo
------------- | -------------
  gorilla/mux | https://github.com/gorilla/mux

## endpoints
	Index  	       	          ,method = GET  	,path = "/"
	GetBlockchain 		  ,method = GET  	,path = "/blockchain"
	RegisterAndBroadcastLoan  ,method = POST  	,path = "/loan/broadcast"
	RegisterAndBroadcastNode  ,method = POST  	,path = "/register-and-broadcast-node"( the new node url => `json:"newnodeurl"`)
	GetLoansForUser  	  ,method = GET  	,path = /user/{userName}
	Mine  			  ,method = GET  	,path = "/mine"
	Consensus  		  ,method = GET  	,path = "/consensus"
		


