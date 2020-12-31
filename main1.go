package main

import (
	"fmt"

	"github.com/TibebeJs/yenepay.sdk.go/checkout"
)

func main() {

	yenepay := checkout.NewYenePayCheckOut()

	fmt.Println("***************[ Express Checkout ]***************")
	fmt.Println(yenepay.ExpressCheckoutUrl(checkout.NewCheckoutOption(
		true,
		checkout.ExpressCheckout,
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
	), checkout.NewExpressCheckoutItem(
		"544",
		"PC",
		30.0,
		2,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	)))

	fmt.Println("***************[ Cart Checkout ]***************")

	fmt.Println(yenepay.CartCheckoutUrl(checkout.NewCheckoutOption(
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
	), []checkout.CartCheckoutItem{
		{
			ItemId:    "544",
			ItemName:  "PC",
			UnitPrice: 30.0,
			Quantity:  2,
		},
		{
			ItemId:    "541",
			ItemName:  "PC2",
			UnitPrice: 30.0,
			Quantity:  2,
		},
	}))
}
