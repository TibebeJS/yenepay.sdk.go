package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCheckoutOption(t *testing.T) {
	actual := NewOption(
		OptionsParams{
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
		},
	)

	expected := &Option{
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

// func TestOptionsToJSONForExpress(t *testing.T) {
// 	actual, _ := NewOption(
// 		OptionsParams{
// 			true,
// 			ExpressCheckout,
// 			"2",
// 			"localhost:8000/success",
// 			"localhost:8000/cancel",
// 			"localhost:8000/ipn",
// 			"localhost:8000/failure",
// 			2,
// 			"2",
// 			10.0,
// 			2.0,
// 			0.0,
// 			0.0,
// 			0.0,
// 		},
// 	).ToJSON(false)

// 	expected := `
// 	{
// 		"UseSandbox": true,
// 		"Process": "Express",
// 		"MerchantID": "2",
// 		"SuccessURL": "localhost:8000/success",
// 		"CancelURL": "localhost:8000/cancel",
// 		"IPNURL": "localhost:8000/ipn",
// 		"FailureURL": "localhost:8000/failure",
// 		"ExpiresAfter": 2,
// 		"MerchantOrderID": "2"
// 	}
// 	`

// 	require.JSONEq(t, expected, actual)

// }
// func TestOptionsToJSONForCart(t *testing.T) {
// 	actual, _ := NewOption(
// 		OptionsParams{
// 			true,
// 			ExpressCheckout,
// 			"2",
// 			"localhost:8000/success",
// 			"localhost:8000/cancel",
// 			"localhost:8000/ipn",
// 			"localhost:8000/failure",
// 			2,
// 			"2",
// 			10.0,
// 			2.0,
// 			0.0,
// 			0.0,
// 			0.0,
// 		},
// 	).ToJSON(true)

// 	expected := `
// 	{
// 		"UseSandbox": true,
// 		"Process": "Express",
// 		"MerchantID": "2",
// 		"SuccessURL": "localhost:8000/success",
// 		"CancelURL": "localhost:8000/cancel",
// 		"IPNURL": "localhost:8000/ipn",
// 		"FailureURL": "localhost:8000/failure",
// 		"ExpiresAfter": 2,
// 		"MerchantOrderID": "2",
// 		"TotalItemsDeliveryFee": 10.0,
// 		"TotalItemsTax1": 2.0,
// 		"TotalItemsTax2": 0.0,
// 		"TotalItemsDiscount": 0.0,
// 		"TotalItemsHandlingFee": 0.0
// 	}
// 	`

// 	require.JSONEq(t, expected, actual)

// }
