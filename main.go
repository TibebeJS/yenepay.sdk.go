package main

import (
	"fmt"

	"github.com/TibebeJs/yenepay.sdk.go/checkout"
)

func main() {
	checkoutOption := checkout.NewCheckoutOption(
		true,
		checkout.ExpressCheckout,
		"222",
		"sdfsd",
		"sdfsd",
		"sdfsd",
		"sdfsd",
		3,
		"sdfsdsdfsd",
		2.0,
		3.0,
		1.0,
		5.0,
		10.0,
	)

	checkoutItem := checkout.NewCheckoutItem(
		"544",
		"PC",
		30.0,
		2,
		0.2,
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
}
