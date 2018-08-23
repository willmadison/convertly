package convertly_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/convertly"
)

func TestFixerExchanger(t *testing.T) {
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

	e := convertly.NewFixerExchanger("43d8eaede49bad82ef3e2c4dbcfafbbe")

	for _, c := range cases {
		t.Run(fmt.Sprintf("converting %v %v to %v", c.amount, c.from, c.to), func(t *testing.T) {
			actual, err := e.Exchange(c.from, c.to, c.amount)
			assert.Nil(t, err)
			assert.NotNil(t, actual)
			assert.Equal(t, c.expected.Denomination, actual.Denomination)
			assert.True(t, actual.Amount > 0, "should be properly converted to a non zero transaction")
		})
	}
}
