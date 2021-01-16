package main

import (
	"fmt"

	"github.com/TibebeJs/yenepay.sdk.go/checkout"
)

func main() {

	yenepay := checkout.NewYenePayCheckOut()

	fmt.Println("***************[ Express Checkout ]***************")
	fmt.Println(yenepay.ExpressCheckoutURL(checkout.NewCheckoutOption(
		checkout.OptionsParams{
			UseSandbox:      true,
			Process:         checkout.ExpressCheckout,
			MerchantID:      "0694",
			SuccessURL:      "http://localhost:3000/Home/PaymentSuccessReturnURL",
			CancelURL:       "http://localhost:3000/Home/CancelReturnURL",
			IPNURL:          "http://localhost:3000/Home/IPNURL",
			FailureURL:      "http://localhost:3000/Home/FailureURL",
			ExpiresAfter:    2880,
			MerchantOrderID: "ab-cd",
		},
	), checkout.NewExpressCheckoutItem(
		checkout.ExpressParams{ItemID: "544", ItemName: "PC", UnitPrice: 30.0, Quantity: 2},
	)))

	fmt.Println("***************[ Cart Checkout ]***************")

	fmt.Println(yenepay.CartCheckoutURL(checkout.NewCheckoutOption(
		checkout.OptionsParams{
			UseSandbox:      true,
			Process:         checkout.ExpressCheckout,
			MerchantID:      "0694",
			SuccessURL:      "http://localhost:3000/Home/PaymentSuccessReturnURL",
			CancelURL:       "http://localhost:3000/Home/CancelReturnURL",
			IPNURL:          "http://localhost:3000/Home/IPNURL",
			FailureURL:      "http://localhost:3000/Home/FailureURL",
			ExpiresAfter:    2880,
			MerchantOrderID: "ab-cd",
		},
	), []checkout.CartCheckoutItem{
		{
			ItemID:    "544",
			ItemName:  "PC",
			UnitPrice: 30.0,
			Quantity:  2,
		},
		{
			ItemID:    "541",
			ItemName:  "PC2",
			UnitPrice: 30.0,
			Quantity:  2,
		},
	}))
}
