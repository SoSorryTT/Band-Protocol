package problem3

import (
	"assignment/band-protocol/problem3/pkg"
	"log"
	"time"
)

func Problem3() {
	server := pkg.NewServerClient()

	// 1. Broadcast Transaction
	// Create a payload
	payload := pkg.SentTransactionPayload{
		Symbol:    "ETH",
		Price:     4500,
		Timestamp: uint64(time.Now().Unix()),
	}
	// Create a http request
	result, err := server.SentTransaction(payload)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("1. Broadcast Transaction")
	log.Println("tx_hash :", result.TxHash)

	// 2. Transaction Status Monitoring
	// Create a http request
	result2, err := server.GetTransactionStatus(result.TxHash)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("2. Transaction Status Monitoring")
	log.Println("tx_status :", result2.TxStatus)
}
