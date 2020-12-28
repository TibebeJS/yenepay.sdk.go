package checkout

import (
	"encoding/json"
)

type CheckoutItem struct {
	ItemId      string
	ItemName    string
	UnitPrice   float64
	Quantity    int
	Discount    float64
	HandlingFee float64
	DeliveryFee float64
	Tax1        float64
	Tax2        float64
}

func (self *CheckoutItem) ToJSON() (string, error) {
	data, err := json.Marshal(self)

	if err != nil {
		return "", err
	}
	return string(data), nil
}

func NewCheckoutItem(ItemId string, ItemName string, UnitPrice float64, Quantity int, Tax1 float64, Tax2 float64, Discount float64, HandlingFee float64, DeliveryFee float64) *CheckoutItem {
	return &CheckoutItem{
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
