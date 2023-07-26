package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/murilojssilva/go_planejamento_financeiro/model/transaction"
	"github.com/murilojssilva/go_planejamento_financeiro/util"
)

// GetTransactions show all registered transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) { // visibilidade privada por começar com letra minúscula
	if r.Method != "GET" {
		w.WriteHeader(http.StatusHTTPVersionNotSupported)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:           "Salário",
			Amount:          1200,
			TransactionType: 0,
			CreatedAt:       util.StringToTime("2023-03-05T10:10:10"),
		},
	}
	_ = json.NewEncoder(w).Encode(transactions)
}

// CreateTransactions create a new transaction
func CreateTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusHTTPVersionNotSupported)
		return
	}

	var res = transaction.Transactions{}

	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Println(res)
}
