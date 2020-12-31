/*
Package checkout implements all checkout related logic and includes all checkout related models.

Models included in this package:

    PDT: Payment Data Transfer
    IPN: Instant Payment Notification
    Checkout: YenePay's Payment Order
*/
package checkout

import (
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
func (self *YenePayCheckOut) GetCheckoutUrlForCart(checkoutOptions *CheckoutOption, checkoutItems []*CheckoutItem) string {

	// https://test.yenepay.com/Home/Process/
	// ?Process=Cart&
	// MerchantId=YOUR_YENEPAY_SELLER_CODE&
	// SuccessUrl=http://localhost:3000/Home/PaymentSuccessReturnUrl&CancelUrl=
	// &IPNUrl=http://localhost:3000/Home/IPNDestination&FailureUrl=&
	// ExpiresAfter=2880&MerchantOrderId=ab-cd
	// &TotalItemsDeliveryFee=100&
	// TotalItemsTax1=345&TotalItemsTax2=0&TotalItemsDiscount=50&TotalItemsHandlingFee=30&
	// Items[0].ItemId=2&Items[0].ItemName=Watch&Items[0].UnitPrice=2000&Items[0].Quantity=1&Items[1].ItemId=1&Items[1].ItemName=Bag&Items[1].UnitPrice=300&Items[1].Quantity=1

	v, _ := query.Values(checkoutOptions)

	checkoutUrl := self.CheckoutBaseUrlProd + "?" + v.Encode()

	if checkoutOptions.UseSandbox {
		checkoutUrl = self.CheckoutBaseUrlSandbox + "?" + v.Encode()
	}

	return checkoutUrl
}
