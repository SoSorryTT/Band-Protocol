package pkg

import (
	"encoding/json"
	"fmt"
)

type SentTransactionPayload struct {
	Symbol    string
	Price     uint64
	Timestamp uint64
}

type SentTransactionResponseBody struct {
	TxHash string `json:"tx_hash"`
}

type Transaction struct {
	TxHash string
}

func (c ServerClient) SentTransaction(input SentTransactionPayload) (Transaction, error) {
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

	resBody, err := c.Execute(endpoint, "POST", jsonResultQuery)
	if err != nil {
		return Transaction{}, err
	}

	var response SentTransactionResponseBody
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return Transaction{}, err
	}

	return Transaction{
		TxHash: response.TxHash,
	}, nil
}
