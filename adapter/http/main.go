package http

import (
	"net/http"

	"github.com/murilojssilva/go_planejamento_financeiro/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransactions)
	http.ListenAndServe(":8080", nil)
}