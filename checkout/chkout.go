package checkout

import (
	"github.com/google/go-querystring/query"
)

type YenePayCheckOut struct {
	checkoutBaseUrlProd    string
	checkoutBaseUrlSandbox string
	ipnVerifyUrlProd       string
	ipnVerifyUrlSandbox    string
	pdtUrlProd             string
	pdtUrlSandbox          string
}

func NewYenePayCheckOut() *YenePayCheckOut {
	self := &YenePayCheckOut{}

	self.checkoutBaseUrlProd = "https://www.yenepay.com/checkout/Home/Process/"
	self.checkoutBaseUrlSandbox = "https://test.yenepay.com/Home/Process/"
	self.ipnVerifyUrlProd = "https://endpoints.yenepay.com/api/verify/ipn/"
	self.ipnVerifyUrlSandbox = "https://testapi.yenepay.com/api/verify/ipn/"
	self.pdtUrlProd = "https://endpoints.yenepay.com/api/verify/pdt/"
	self.pdtUrlSandbox = "https://testapi.yenepay.com/api/verify/pdt/"

	return self
}

func (self *YenePayCheckOut) GetCheckoutUrlForExpress(checkoutOptions *CheckoutOption, checkoutItem *CheckoutItem) string {

	v, _ := query.Values(checkoutOptions.GetCartFields())

	checkoutUrl := self.checkoutBaseUrlProd + "?" + v.Encode()

	if checkoutOptions.UseSandbox {
		checkoutUrl = self.checkoutBaseUrlSandbox + "?" + v.Encode()
	}

	return checkoutUrl
}
