package pkg

import "encoding/json"

// Struct receive a body from server
type GetTransactionStatusResponseBody struct {
	TxStatus string `json:"tx_status"`
}

// Internal struct for return
type TransactionStatus struct {
	TxStatus string
}

func (c ServerClient) GetTransactionStatus(input string) (TransactionStatus, error) {
	endpoint := "check/" + input

	//Sent a http request
	resBody, err := c.Execute(endpoint, "GET", nil)
	if err != nil {
		return TransactionStatus{}, err
	}

	//Receive a response body
	var response GetTransactionStatusResponseBody
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return TransactionStatus{}, err
	}

	return TransactionStatus{
		TxStatus: response.TxStatus,
	}, nil
}
