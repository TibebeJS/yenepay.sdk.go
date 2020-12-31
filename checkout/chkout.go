/*
Package checkout implements all checkout related logic and includes all checkout related models.

Models included in this package:

    PDT: Payment Data Transfer
    IPN: Instant Payment Notification
    Checkout: YenePay's Payment Order
*/
package checkout

import (
	"fmt"
	"reflect"

	"github.com/google/go-querystring/query"
)

type YenePayCheckOut struct {
	CheckoutBaseUrlProd    string
	CheckoutBaseUrlSandbox string
	IpnVerifyUrlProd       string
	IpnVerifyUrlSandbox    string
	PdtUrlProd             string
	PdtUrlSandbox          string
}

// YenePayCheckOut Constructor
func NewYenePayCheckOut() *YenePayCheckOut {
	self := &YenePayCheckOut{}

	self.CheckoutBaseUrlProd = "https://www.yenepay.com/checkout/Home/Process/"
	self.CheckoutBaseUrlSandbox = "https://test.yenepay.com/Home/Process/"
	self.IpnVerifyUrlProd = "https://endpoints.yenepay.com/api/verify/ipn/"
	self.IpnVerifyUrlSandbox = "https://testapi.yenepay.com/api/verify/ipn/"
	self.PdtUrlProd = "https://endpoints.yenepay.com/api/verify/pdt/"
	self.PdtUrlSandbox = "https://testapi.yenepay.com/api/verify/pdt/"

	return self
}

// Generate Checkout URL for Express Checkout
func (self *YenePayCheckOut) GetCheckoutUrlForExpress(checkoutOptions *CheckoutOption, checkoutItem *CheckoutItem) string {

	v, _ := query.Values(checkoutOptions.GetExpressFields())

	checkoutUrl := self.CheckoutBaseUrlProd + "?" + v.Encode()

	if checkoutOptions.UseSandbox {
		checkoutUrl = self.CheckoutBaseUrlSandbox + "?" + v.Encode()
	}

	return checkoutUrl
}

// Generate Checkout URL for Cart Checkout
func (self *YenePayCheckOut) GetCheckoutUrlForCart(checkoutOptions *CheckoutOption, checkoutItems []CheckoutItem) string {

	v, _ := query.Values(checkoutOptions)

	checkoutUrl := self.CheckoutBaseUrlProd + "?" + v.Encode()

	if checkoutOptions.UseSandbox {
		checkoutUrl = self.CheckoutBaseUrlSandbox + "?" + v.Encode()
	}

	for i, item := range checkoutItems {
		v := reflect.ValueOf(item.GetCartFields())
		typeOfS := v.Type()

		for j := 0; j < v.NumField(); j++ {
			checkoutUrl += fmt.Sprintf("&Items[%d].%s=%v", i, typeOfS.Field(j).Name, v.Field(j).Interface())
		}
	}

	return checkoutUrl
}
