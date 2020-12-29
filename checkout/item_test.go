package checkout

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestItemToJSON(t *testing.T) {
	actual, _ := NewCheckoutItem(
		"2",
		"item-name",
		10.0,
		2,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	).ToJSON()

	expected := `
	{
		"ItemId" : "2",
		"ItemName" : "item-name",
		"UnitPrice": 10.0,
		"Quantity": 2,
		"Discount": 0.0,
		"HandlingFee": 0.0,
		"DeliveryFee": 0.0,
		"Tax1": 0.0,
		"Tax2": 0.0
	}
	`

	require.JSONEq(t, expected, actual)

}

func TestNewCheckoutItem(t *testing.T) {
	item := NewCheckoutItem(
		"2",
		"item-name",
		10.0,
		2,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	)

	expected := &CheckoutItem{
		"2",
		"item-name",
		10.0,
		2,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	}

	if !cmp.Equal(item, expected) {
		t.Error("NewCheckoutItem Constructor is not working as expected")
	}
}
