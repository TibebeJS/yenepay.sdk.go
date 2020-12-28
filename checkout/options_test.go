package checkout

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewCheckoutOption(t *testing.T) {
	options := NewCheckoutOption(
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

	if !cmp.Equal(options, expected) {
		t.Error("NewCheckoutOption Constructor is not working as expected")
	}
}
