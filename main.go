package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() { 
	http.HandleFunc("/transactions", getTransactions)
	http.HandleFunc("/transactions/create", createTransactions)
	http.ListenAndServe(":8080", nil)
}

type Transaction struct { // visibilidade pública por começar com letra maiúscula
	Title string `json:"title"`
	Amount float32 `json:"amount"`
	TransactionType int `json:"transaction_type"`
	CreatedAt time.Time `json:"created_at"`
}

type Transactions []Transaction

func getTransactions (w http.ResponseWriter, r *http.Request) { // visibilidade privada por começar com letra minúscula
	if r.Method != "GET" {
		w.WriteHeader(http.StatusHTTPVersionNotSupported)
		return 
	}

	w.Header().Set("Content-type","application/json")
	var layout = "2006-01-02T15:05:05"

		salaryReceived, _ := time.Parse(layout, "2020-04-05T20:10:05")

	var transactions = Transactions {
		Transaction{
			Title: "Salário",
			Amount: 1200,
			TransactionType: 0,
			CreatedAt: salaryReceived,
		},
	}
	_ = json.NewEncoder(w).Encode(transactions)
}

func createTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusHTTPVersionNotSupported)
		return 
	}

	
}