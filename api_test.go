package convertly_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/convertly"
)

type dummyExchange struct{}

func (d dummyExchange) Exchange(from, to convertly.Currency, amount float64) (convertly.Transaction, error) {
	return convertly.Transaction{
		25.0,
		to,
	}, nil
}

func TestConversionPairs(t *testing.T) {
	cases := []struct {
		amount   float64
		from     convertly.Currency
		to       convertly.Currency
		expected convertly.Transaction
	}{
		{
			amount: 1.0,
			from:   convertly.USD,
			to:     convertly.GBP,
			expected: convertly.Transaction{
				Amount:       25.0,
				Denomination: convertly.GBP,
			},
		},
	}

	handler := convertly.RootHandler(dummyExchange{})

	for _, c := range cases {
		t.Run(fmt.Sprintf("converting %v %v to %v", c.amount, c.from, c.to), func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/exchange/%v/%v?amount=%v", c.from, c.to, c.amount), nil)
			w := httptest.NewRecorder()

			handler.ServeHTTP(w, req)

			response := w.Result()
			body, _ := ioutil.ReadAll(response.Body)

			assert.Equal(t, http.StatusOK, response.StatusCode)

			var actual convertly.Transaction

			json.Unmarshal(body, &actual)

			assert.Equal(t, c.expected, actual)
		})
	}

}
