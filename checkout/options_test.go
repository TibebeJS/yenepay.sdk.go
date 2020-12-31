package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCheckoutOption(t *testing.T) {
	actual := NewCheckoutOption(
		true,
		ExpressCheckout,
		"2",
		"localhost:8000/success",
		"localhost:8000/cancel",
		"localhost:8000/ipn",
		"localhost:8000/failure",
		2,
		"2",
		10.0,
		2.0,
		0.0,
		0.0,
		0.0,
	)

	expected := &CheckoutOption{
		true,
		ExpressCheckout,
		"2",
		"localhost:8000/success",
		"localhost:8000/cancel",
		"localhost:8000/ipn",
		"localhost:8000/failure",
		2,
		"2",
		10.0,
		2.0,
		0.0,
		0.0,
		0.0,
	}

	assert.Exactly(t, expected, actual)
}

func TestOptionsToJSONForExpress(t *testing.T) {
	actual, _ := NewCheckoutOption(
		true,
		ExpressCheckout,
		"2",
		"localhost:8000/success",
		"localhost:8000/cancel",
		"localhost:8000/ipn",
		"localhost:8000/failure",
		2,
		"2",
		10.0,
		2.0,
		0.0,
		0.0,
		0.0,
	).ToJSON(false)

	expected := `
	{
		"UseSandbox": true,
		"Process": "Express",
		"MerchantId": "2",
		"SuccessUrl": "localhost:8000/success",
		"CancelUrl": "localhost:8000/cancel",
		"IPNUrl": "localhost:8000/ipn",
		"FailureUrl": "localhost:8000/failure",
		"ExpiresAfter": 2,
		"MerchantOrderId": "2"
	}
	`

	require.JSONEq(t, expected, actual)

}
func TestOptionsToJSONForCart(t *testing.T) {
	actual, _ := NewCheckoutOption(
		true,
		ExpressCheckout,
		"2",
		"localhost:8000/success",
		"localhost:8000/cancel",
		"localhost:8000/ipn",
		"localhost:8000/failure",
		2,
		"2",
		10.0,
		2.0,
		0.0,
		0.0,
		0.0,
	).ToJSON(true)

	expected := `
	{
		"UseSandbox": true,
		"Process": "Express",
		"MerchantId": "2",
		"SuccessUrl": "localhost:8000/success",
		"CancelUrl": "localhost:8000/cancel",
		"IPNUrl": "localhost:8000/ipn",
		"FailureUrl": "localhost:8000/failure",
		"ExpiresAfter": 2,
		"MerchantOrderId": "2",
		"TotalItemsDeliveryFee": 10.0,
		"TotalItemsTax1": 2.0,
		"TotalItemsTax2": 0.0,
		"TotalItemsDiscount": 0.0,
		"TotalItemsHandlingFee": 0.0
	}
	`

	require.JSONEq(t, expected, actual)

}
