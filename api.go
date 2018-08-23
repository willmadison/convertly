package convertly

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Transaction struct {
	Amount       float64
	Denomination Currency
}

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
)

func RootHandler(e Exchanger) http.Handler {
	r := mux.NewRouter()

	//r.HandleFunc("/exchange/{from}/{to}", func(w http.ResponseWriter, r *http.Request) {})

	return r
}
