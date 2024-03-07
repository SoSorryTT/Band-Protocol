package pkg

import (
	"bytes"
	"io"
	"net/http"
)

type ServerClient struct {
	//key or something for auth here
	//usually it will be an auth to put in header
	//but in this case I think the question did not provide,
	//so I'm going to skip
}

func NewServerClient() ServerClient {
	return ServerClient{}
}

type TransactionPayload struct {
	sent SentTransactionPayload
}

func (c ServerClient) Execute(endpointRoute string, requestMethod string, bodyPayload []byte) ([]byte, error) {
	endpoint := "https://mock-node-wgqbnxruha-as.a.run.app/" + endpointRoute

	resp, err := http.NewRequest(requestMethod, endpoint, bytes.NewBuffer(bodyPayload))
	if err != nil {
		return nil, err
	}
	resp.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	response, err := client.Do(resp)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
