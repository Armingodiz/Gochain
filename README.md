# Gochain
simple bank system for holding loans using blockchain 


## endpoints
	Index  	       	          ,method = GET  	,path = "/"
	GetBlockchain 		  ,method = GET  	,path = "/blockchain"
	RegisterAndBroadcastNode  ,method = POST  	,path = "/register-and-broadcast-node"( the new node url => `json:"newnodeurl"`)
	GetLoansForUser  	  ,method = GET  	,path = /user/{userName}
	Mine  			  ,method = GET  	,path = "/mine"
	Consensus  		  ,method = GET  	,path = "/consensus"
		


