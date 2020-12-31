package checkout

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestItemToJSON(t *testing.T) {
	actual, _ := NewExpressCheckoutItem(
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

	_, actualError := (&ExpressCheckoutItem{}).ToJSON()

	if assert.Error(t, actualError) {
		for _, err := range actualError.(validator.ValidationErrors) {
			switch err.Field() {
			case "ItemId":
				assert.Exactly(t, "ItemId", err.Field())
				assert.Exactly(t, "required", err.Tag())
			case "ItemName":
				assert.Exactly(t, "ItemName", err.Field())
				assert.Exactly(t, "required", err.Tag())
			}

		}
	}
}

func TestNewExpressCheckoutItem(t *testing.T) {
	item := NewExpressCheckoutItem(
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

	expected := &ExpressCheckoutItem{
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

	assert.Exactly(t, expected, item)
}
func TestNewCartCheckoutItem(t *testing.T) {
	item := NewCartCheckoutItem(
		"2",
		"item-name",
		10.0,
		2,
	)

	expected := &CartCheckoutItem{
		"2",
		"item-name",
		10.0,
		2,
	}

	assert.Exactly(t, expected, item)
}
