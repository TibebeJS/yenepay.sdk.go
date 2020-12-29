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

func (self *YenePayCheckOut) GetCheckoutUrlForExpress(checkoutOptions *CheckoutOption, checkoutItem *CheckoutItem) string {

	v, _ := query.Values(checkoutOptions.GetCartFields())

	checkoutUrl := self.CheckoutBaseUrlProd + "?" + v.Encode()

	if checkoutOptions.UseSandbox {
		checkoutUrl = self.CheckoutBaseUrlSandbox + "?" + v.Encode()
	}

	return checkoutUrl
}
