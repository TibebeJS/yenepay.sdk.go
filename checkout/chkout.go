/*
Package checkout implements all checkout related logic and includes all checkout related models.

Models included in this package:

    PDT: Payment Data Transfer
    IPN: Instant Payment Notification
    Checkout: YenePay's Payment Order
*/
package checkout

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
func (self *YenePayCheckOut) ExpressCheckoutUrl(checkoutOptions *CheckoutOption, checkoutItem *ExpressCheckoutItem) string {

	optsQs, _ := query.Values(checkoutOptions.GetExpressFields())

	itemQs, _ := query.Values(checkoutItem)

	checkoutUrl := self.CheckoutBaseUrlProd + "?" + optsQs.Encode() + "&" + itemQs.Encode()

	if checkoutOptions.UseSandbox {
		checkoutUrl = self.CheckoutBaseUrlSandbox + "?" + optsQs.Encode() + "&" + itemQs.Encode()
	}

	return checkoutUrl
}

// Generate Checkout URL for Cart Checkout
func (self *YenePayCheckOut) CartCheckoutUrl(checkoutOptions *CheckoutOption, cartItems []CartCheckoutItem) string {

	v, _ := query.Values(checkoutOptions)

	checkoutUrl := self.CheckoutBaseUrlProd + "?" + v.Encode()

	if checkoutOptions.UseSandbox {
		checkoutUrl = self.CheckoutBaseUrlSandbox + "?" + v.Encode()
	}

	for i, item := range cartItems {
		v := reflect.ValueOf(item)
		typeOfS := v.Type()

		for j := 0; j < v.NumField(); j++ {
			checkoutUrl += fmt.Sprintf("&Items[%d].%s=%v", i, typeOfS.Field(j).Name, v.Field(j).Interface())
		}
	}

	return checkoutUrl
}

// WIP: check IPN authenticity
func (self *YenePayCheckOut) IsIPNAuthentic(ipnModel interface{}, useSandbox bool) (bool, error) {
	ipnUrl := self.IpnVerifyUrlProd

	if useSandbox {
		ipnUrl = self.IpnVerifyUrlSandbox
	}

	reqBody, _ := json.Marshal(ipnModel)
	resp, err := http.Post(ipnUrl, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {

	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

	}

	fmt.Println(body)

	return true, nil
}

// WIP: Request PDT model
func (self *YenePayCheckOut) RequestPDT(pdtReq PdtRequestModel) (interface{}, error) {
	pdtUrl := self.PdtUrlProd

	if pdtReq.UseSandbox {
		pdtUrl = self.PdtUrlSandbox
	}

	reqBody, _ := json.Marshal(pdtReq)
	resp, err := http.Post(pdtUrl, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {

	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
