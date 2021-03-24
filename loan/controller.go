package loan

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

/////////////////////////////////////////////////////////////////// Making calls for broadcasting part
//MakeCall :
func MakeCall(mode string, url string, jsonStr []byte) interface{} {
	// call url in node
	log.Println(mode)
	log.Println(url)
	req, err := http.NewRequest(mode, url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in call " + url)
		log.Println(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	var returnValue interface{}
	if err := json.Unmarshal(respBody, &returnValue); err != nil { // unmarshal body contents as a type Candidate
		if err != nil {
			log.Fatalln("Error "+url+" unmarshalling data", err)
			return nil
		}
	}
	log.Println(returnValue)
	return returnValue
}

//MakePostCall
func MakePostCall(url string, jsonStr []byte) {
	// call url in POST
	MakeCall("POST", url, jsonStr)
}

//MakeGetCall
func MakeGetCall(url string, jsonStr []byte) interface{} {
	// call url in GET
	return MakeCall("GET", url, jsonStr)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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
///////////////////////////////////////////////////////////////////////////////////////////// mining part :
//Mine GET /mine
func (c *Controller) Mine(w http.ResponseWriter, r *http.Request) {
	lastBlock := c.blockchain.GetLastBlock()
	previousBlockHash := lastBlock.Hash
	currentBlockData := BlockData{Index: strconv.Itoa(lastBlock.Index - 1), Loans: c.blockchain.PendingLoans}
	currentBlockDataAsByteArray, _ := json.Marshal(currentBlockData)
	currentBlockDataAsStr := base64.URLEncoding.EncodeToString(currentBlockDataAsByteArray)

	nonce := c.blockchain.ProofOfWork(previousBlockHash, currentBlockDataAsStr)
	blockHash := c.blockchain.HashBlock(previousBlockHash, currentBlockDataAsStr, nonce)
	newBlock := c.blockchain.CreateNewBlock(nonce, previousBlockHash, blockHash)
	blockToBroadcast, _ := json.Marshal(newBlock)

	for _, node := range c.blockchain.NetworkNodes {
		if node != c.currentNodeURL {
			// call /receive-new-block in node
			MakePostCall(node+"/receive-new-block", blockToBroadcast)
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var resp ResponseToSend
	resp.Note = "New block mined and broadcast successfully."
	data, _ := json.Marshal(resp)
	w.Write(data)
	return
}

//ReceiveNewBlock POST /receive-new-block
func (c *Controller) ReceiveNewBlock(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body) // read the body of the request
	if err != nil {
		log.Fatalln("Error ReceiveNewBlock", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error ReceiveNewBlock", err)
	}

	var blockReceived Block
	if err := json.Unmarshal(body, &blockReceived); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error ReceiveNewBlock unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	var resp ResponseToSend
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// append block to blockchain
	if c.blockchain.CheckNewBlockHash(blockReceived) {
		resp.Note = "New Block received and accepted."
		c.blockchain.PendingLoans = Loans{}
		c.blockchain.Chain = append(c.blockchain.Chain, blockReceived)
	} else {
		resp.Note = "New Block rejected."
	}

	data, _ := json.Marshal(resp)
	w.Write(data)
	return
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
