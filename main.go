package main

import (
	"fmt"

	"github.com/TibebeJs/yenepay.sdk.go/checkout"
)

func main() {
	checkoutOption := checkout.NewCheckoutOption(
		true,
		checkout.CartCheckout,
		"0694",
		"http://localhost:3000/Home/PaymentSuccessReturnUrl",
		"http://localhost:3000/Home/CancelReturnUrl",
		"http://localhost:3000/Home/IPNUrl",
		"http://localhost:3000/Home/FailureUrl",
		2880,
		"ab-cd",
		2.0,
		3.0,
		0.0,
		5.0,
		10.0,
	)

	checkoutItem := checkout.NewCheckoutItem(
		"544",
		"PC",
		30.0,
		2,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	)

	json, _ := checkoutItem.ToJSON()

	fmt.Println(json)
	fmt.Println(checkoutOption.ToJSON(false))

	check := checkout.NewYenePayCheckOut()

	fmt.Println("***************************")
	fmt.Println(check.GetCheckoutUrlForExpress(checkoutOption, checkoutItem))

	fmt.Println("***************************")

	fmt.Println(check.GetCheckoutUrlForCart(checkoutOption, []checkout.CheckoutItem{
		{
			ItemId:      "544",
			ItemName:    "PC",
			UnitPrice:   30.0,
			Quantity:    2,
			Discount:    0.0,
			HandlingFee: 0.0,
			DeliveryFee: 0.0,
			Tax1:        0.0,
			Tax2:        0.0,
		},
		{
			ItemId:      "541",
			ItemName:    "PC2",
			UnitPrice:   30.0,
			Quantity:    2,
			Discount:    0.0,
			HandlingFee: 0.0,
			DeliveryFee: 0.0,
			Tax1:        0.0,
			Tax2:        0.0,
		},
	}))
}
