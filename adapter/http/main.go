package http

import (
	"net/http"

	"github.com/murilojssilva/go_planejamento_financeiro/adapter/http/actuator"
	"github.com/murilojssilva/go_planejamento_financeiro/adapter/http/transaction"
)

// Init have the project routes
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransactions)

	http.HandleFunc("/healthy", actuator.Healthy)

	http.ListenAndServe(":8080", nil)
}
