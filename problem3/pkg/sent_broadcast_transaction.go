package pkg

import (
	"encoding/json"
	"fmt"
)

// SentTransactionPayload Struct of payload
type SentTransactionPayload struct {
	Symbol    string
	Price     uint64
	Timestamp uint64
}

// SentTransactionResponseBody Struct receive a body from server
type SentTransactionResponseBody struct {
	TxHash string `json:"tx_hash"`
}

// Transaction Internal struct for return
type Transaction struct {
	TxHash string
}

// SentTransaction run time is O(n) because it depend on return payload data
func (c ServerClient) SentTransaction(input SentTransactionPayload) (Transaction, error) {
	// Convert and marshaling a payload input to JSON
	jsonMapInstanceQuery := map[string]interface{}{
		"symbol":    input.Symbol,
		"price":     input.Price,
		"timestamp": input.Timestamp,
	}
	jsonResultQuery, err := json.Marshal(jsonMapInstanceQuery)
	if err != nil {
		fmt.Printf("There was an error marshaling the JSON instance %v", err)
	}

	endpoint := "broadcast"

	//Sent a http request
	resBody, err := c.Execute(endpoint, "POST", jsonResultQuery)
	if err != nil {
		return Transaction{}, err
	}

	//Receive a response body
	var response SentTransactionResponseBody
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return Transaction{}, err
	}

	return Transaction{
		TxHash: response.TxHash,
	}, nil
}
