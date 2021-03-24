package loan

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Controller :
type Controller struct {
	blockchain     *Blockchain
	currentNodeURL string
}

//ResponseToSend ...
type ResponseToSend struct {
	Note string
}

//Index
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}

//GetBlockchain
func (c *Controller) GetBlockchain(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(c.blockchain)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

///////////////////////////////////////////////////////////////////////////////////////// Adding new loan :

//RegisterAndBroadcastBet POST /loan/broadcast
func (c *Controller) RegisterAndBroadcastLoan(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body) // read the body of the request
	errMessage := "Error RegisterLoan"
	if err != nil {
		log.Fatalln(errMessage, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln(errMessage, err)
	}
	var loan Loan
	if err := json.Unmarshal(body, &loan); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln(errMessage+" unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.blockchain.RegisterLoan(loan) // registers the bet into the blockchain
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// broadcast => sending post request to add New loan to other nodes too
	for _, node := range c.blockchain.NetworkNodes {
		if node != c.currentNodeURL {
			// call /register-node in node
			MakePostCall(node+"/bet", body)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	var resp ResponseToSend
	resp.Note = "Loan created and broadcast successfully."
	data, _ := json.Marshal(resp)
	w.Write(data)
}

//RegisterBet POST /bet
func (c *Controller) RegisterLoan(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body) // read the body of the request
	errMessage := "Error RegisterLoan"
	if err != nil {
		log.Fatalln(errMessage, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln(errMessage, err)
	}
	var loan Loan
	if err := json.Unmarshal(body, &loan); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln(errMessage+" unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.blockchain.RegisterLoan(loan) // registers the bet into the blockchain
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	var resp ResponseToSend
	resp.Note = "Loan created and broadcast successfully."
	data, _ := json.Marshal(resp)
	w.Write(data)
	return
}

////////////////////////////////////////////////////////////////////////////////////////// end loan registration part
