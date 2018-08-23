package convertly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Exchanger interface {
	Exchange(from, to Currency, amount float64) (Transaction, error)
}

type fixerExchanger struct {
	apiToken string
}

type fixerResponse struct {
	Success bool                 `json:"success"`
	Base    string               `json:"base"`
	Date    string               `json:"date"`
	Rates   map[Currency]float64 `json:"rates"`
}

func (f *fixerExchanger) Exchange(from, to Currency, amount float64) (Transaction, error) {
	var t Transaction

	response, err := http.Get(fmt.Sprintf("http://data.fixer.io/api/latest?access_key=%v", f.apiToken))
	if err != nil {
		return t, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return t, err
	}

	var fr fixerResponse

	err = json.Unmarshal(body, &fr)
	if err != nil {
		return t, err
	}

	t.Denomination = to

	euros := amount * fr.Rates[from]
	t.Amount = euros * fr.Rates[to]

	return t, nil
}

func NewFixerExchanger(apiToken string) Exchanger {
	return &fixerExchanger{apiToken}
}
