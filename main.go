package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() { 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá, bem-vindo a minha página!")
	})
	http.ListenAndServe(":8080", nil)
	fmt.Println("Hello world")
}

type Transaction struct { // visibilidade pública por começar com letra maiúscula
	title string
	amount float32
	transactionType int
	createdAt time.Time
}

type Transactions []Transaction

func getTransactions (w http.ResponseWriter, r *http.Request) { // visibilidade privada por começar com letra minúscula
	if r.Method != "GET" {
		w.WriteHeader(http.StatusHTTPVersionNotSupported)
		return 
	}
	var layout = "2006-01-02T15:05:05"

		salaryReceived, _ := time.Parse(layout, "2020-04-05T20:10:05")

	var transactions = Transactions {
		Transaction{
			title: "Salário",
			amount: 1200,
			transactionType: 0,
			createdAt: salaryReceived,
		},
	}
	_ = json.NewEncoder(w).Encode(transactions)
}