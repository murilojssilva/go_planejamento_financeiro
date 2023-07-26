package http

import (
	"net/http"

	"github.com/murilojssilva/go_planejamento_financeiro/adapter/http/actuator"
	"github.com/murilojssilva/go_planejamento_financeiro/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init have the project routes
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransactions)

	http.HandleFunc("/healthy", actuator.Healthy)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
