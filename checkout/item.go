package checkout

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type ExpressCheckoutItem struct {
	ItemId      string  `validate:"required,min=1"`
	ItemName    string  `validate:"required,min=1"`
	UnitPrice   float64 `validate:"min=0.1"`
	Quantity    int     `validate:"min=1"`
	Discount    float64 `validate:"min=0.0"`
	HandlingFee float64 `validate:"min=0.0"`
	DeliveryFee float64 `validate:"min=0.0"`
	Tax1        float64 `validate:"min=0.0"`
	Tax2        float64 `validate:"min=0.0"`
}

// Validate and Marshal CheckoutItem to JSON format
func (self *ExpressCheckoutItem) ToJSON() (string, error) {
	var validate *validator.Validate = validator.New()

	err := validate.Struct(self)

	if err != nil {
		return "", err
	} else {
		data, _ := json.Marshal(self)

		return string(data), nil
	}
}

type ExpressParams ExpressCheckoutItem

// ExpressCheckoutItem Constructor
func NewExpressCheckoutItem(
	params ExpressParams,
) *ExpressCheckoutItem {
	return &ExpressCheckoutItem{
		params.ItemId,
		params.ItemName,
		params.UnitPrice,
		params.Quantity,
		params.Discount,
		params.HandlingFee,
		params.DeliveryFee,
		params.Tax1,
		params.Tax2,
	}
}

type CartCheckoutItem struct {
	ItemId    string
	ItemName  string
	UnitPrice float64
	Quantity  int
}

type CartParams CartCheckoutItem

// CartCheckoutItem Constructor
func NewCartCheckoutItem(params CartParams) *CartCheckoutItem {
	return &CartCheckoutItem{
		params.ItemId,
		params.ItemName,
		params.UnitPrice,
		params.Quantity,
	}
}
