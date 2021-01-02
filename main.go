package main

import (
	"fmt"

	"github.com/TibebeJs/yenepay.sdk.go/checkout"
)

func main() {

	yenepay := checkout.NewYenePayCheckOut()

	fmt.Println("***************[ Express Checkout ]***************")
	fmt.Println(yenepay.ExpressCheckoutUrl(checkout.NewCheckoutOption(
		checkout.OptionsParams{
			UseSandbox:      true,
			Process:         checkout.ExpressCheckout,
			MerchantId:      "0694",
			SuccessUrl:      "http://localhost:3000/Home/PaymentSuccessReturnUrl",
			CancelUrl:       "http://localhost:3000/Home/CancelReturnUrl",
			IPNUrl:          "http://localhost:3000/Home/IPNUrl",
			FailureUrl:      "http://localhost:3000/Home/FailureUrl",
			ExpiresAfter:    2880,
			MerchantOrderId: "ab-cd",
		},
	), checkout.NewExpressCheckoutItem(
		checkout.ExpressParams{ItemId: "544", ItemName: "PC", UnitPrice: 30.0, Quantity: 2},
	)))

	fmt.Println("***************[ Cart Checkout ]***************")

	fmt.Println(yenepay.CartCheckoutUrl(checkout.NewCheckoutOption(
		checkout.OptionsParams{
			UseSandbox:      true,
			Process:         checkout.ExpressCheckout,
			MerchantId:      "0694",
			SuccessUrl:      "http://localhost:3000/Home/PaymentSuccessReturnUrl",
			CancelUrl:       "http://localhost:3000/Home/CancelReturnUrl",
			IPNUrl:          "http://localhost:3000/Home/IPNUrl",
			FailureUrl:      "http://localhost:3000/Home/FailureUrl",
			ExpiresAfter:    2880,
			MerchantOrderId: "ab-cd",
		},
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
