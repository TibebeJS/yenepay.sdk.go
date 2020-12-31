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

type CartCheckoutItem struct {
	ItemId    string
	ItemName  string
	UnitPrice float64
	Quantity  int
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

// ExpressCheckoutItem Constructor
func NewExpressCheckoutItem(ItemId string, ItemName string, UnitPrice float64, Quantity int, Discount float64, HandlingFee float64, DeliveryFee float64, Tax1 float64, Tax2 float64) *ExpressCheckoutItem {
	return &ExpressCheckoutItem{
		ItemId,
		ItemName,
		UnitPrice,
		Quantity,
		Discount,
		HandlingFee,
		DeliveryFee,
		Tax1,
		Tax2,
	}
}

// CartCheckoutItem Constructor
func NewCartCheckoutItem(ItemId string, ItemName string, UnitPrice float64, Quantity int) *CartCheckoutItem {
	return &CartCheckoutItem{
		ItemId,
		ItemName,
		UnitPrice,
		Quantity,
	}
}
