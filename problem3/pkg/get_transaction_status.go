package pkg

import "encoding/json"

type GetTransactionStatusResponseBody struct {
	TxStatus string `json:"tx_status"`
}

type TransactionStatus struct {
	TxStatus string
}

func (c ServerClient) GetTransactionStatus(input string) (TransactionStatus, error) {
	endpoint := "check/" + input

	resBody, err := c.Execute(endpoint, "GET", nil)
	if err != nil {
		return TransactionStatus{}, err
	}
	var response GetTransactionStatusResponseBody
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return TransactionStatus{}, err
	}

	return TransactionStatus{
		TxStatus: response.TxStatus,
	}, nil
}
