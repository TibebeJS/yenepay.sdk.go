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

const (
	CHECKOUT_BASE_URL_PROD = "https://www.yenepay.com/checkout/Home/Process/"
	IPN_VERIFY_URL_PROD    = "https://endpoints.yenepay.com/api/verify/ipn/"
	PDT_URL_PROD           = "https://endpoints.yenepay.com/api/verify/pdt/"

	CHECKOUT_BASE_URL_SANDBOX = "https://test.yenepay.com/Home/Process/"
	IPN_VERIFY_URL_SANDBOX    = "https://testapi.yenepay.com/api/verify/ipn/"
	PDT_URL_SANDBOX           = "https://testapi.yenepay.com/api/verify/pdt/"
)

type YenePayCheckOut struct {
}

// YenePayCheckOut Constructor
func NewYenePayCheckOut() *YenePayCheckOut {
	return &YenePayCheckOut{}
}

// Generate Checkout URL for Express Checkout
func (self *YenePayCheckOut) ExpressCheckoutUrl(checkoutOptions *CheckoutOption, checkoutItem *ExpressCheckoutItem) string {

	optsQs, _ := query.Values(checkoutOptions.GetExpressFields())

	itemQs, _ := query.Values(checkoutItem)

	checkoutUrl := CHECKOUT_BASE_URL_PROD + "?" + optsQs.Encode() + "&" + itemQs.Encode()

	if checkoutOptions.UseSandbox {
		checkoutUrl = CHECKOUT_BASE_URL_SANDBOX + "?" + optsQs.Encode() + "&" + itemQs.Encode()
	}

	return checkoutUrl
}

// Generate Checkout URL for Cart Checkout
func (self *YenePayCheckOut) CartCheckoutUrl(checkoutOptions *CheckoutOption, cartItems []CartCheckoutItem) string {

	v, _ := query.Values(checkoutOptions)

	checkoutUrl := CHECKOUT_BASE_URL_PROD + "?" + v.Encode()

	if checkoutOptions.UseSandbox {
		checkoutUrl = CHECKOUT_BASE_URL_SANDBOX + "?" + v.Encode()
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
	ipnUrl := IPN_VERIFY_URL_PROD

	if useSandbox {
		ipnUrl = IPN_VERIFY_URL_SANDBOX
	}

	reqBody, _ := json.Marshal(ipnModel)
	resp, err := http.Post(ipnUrl, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		fmt.Print(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(body)

	return true, nil
}

// WIP: Request PDT model
func (self *YenePayCheckOut) RequestPDT(pdtReq PdtRequestModel) (interface{}, error) {
	pdtUrl := PDT_URL_PROD

	if pdtReq.UseSandbox {
		pdtUrl = PDT_URL_SANDBOX
	}

	reqBody, _ := json.Marshal(pdtReq)
	resp, err := http.Post(pdtUrl, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		fmt.Print(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
