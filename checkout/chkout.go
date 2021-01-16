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
	checkoutBaseURLProd string = "https://www.yenepay.com/checkout/Home/Process/"
	ipnVerifyURLProd    string = "https://endpoints.yenepay.com/api/verify/ipn/"
	pdtURLProd          string = "https://endpoints.yenepay.com/api/verify/pdt/"

	checkoutBaseURLSandbox string = "https://test.yenepay.com/Home/Process/"
	ipnVerifyURLSandbox    string = "https://testapi.yenepay.com/api/verify/ipn/"
	pdtURLSandbox          string = "https://testapi.yenepay.com/api/verify/pdt/"
)

// YenePayCheckOut - main checkout model
type YenePayCheckOut struct {
}

// NewYenePayCheckOut - YenePayCheckOut Constructor
func NewYenePayCheckOut() *YenePayCheckOut {
	return &YenePayCheckOut{}
}

// ExpressCheckoutURL - Generate Checkout URL for Express Checkout
func (checkout *YenePayCheckOut) ExpressCheckoutURL(checkoutOptions *Option, checkoutItem *ExpressCheckoutItem) string {

	optsQs, _ := query.Values(checkoutOptions.GetExpressFields())

	itemQs, _ := query.Values(checkoutItem)

	checkoutURL := checkoutBaseURLProd + "?" + optsQs.Encode() + "&" + itemQs.Encode()

	if checkoutOptions.UseSandbox {
		checkoutURL = checkoutBaseURLSandbox + "?" + optsQs.Encode() + "&" + itemQs.Encode()
	}

	return checkoutURL
}

// CartCheckoutURL - Generate Checkout URL for Cart Checkout
func (checkout *YenePayCheckOut) CartCheckoutURL(checkoutOptions *Option, cartItems []CartCheckoutItem) string {

	v, _ := query.Values(checkoutOptions)

	checkoutURL := checkoutBaseURLProd + "?" + v.Encode()

	if checkoutOptions.UseSandbox {
		checkoutURL = checkoutBaseURLSandbox + "?" + v.Encode()
	}

	for i, item := range cartItems {
		v := reflect.ValueOf(item)
		typeOfS := v.Type()

		for j := 0; j < v.NumField(); j++ {
			checkoutURL += fmt.Sprintf("&Items[%d].%s=%v", i, typeOfS.Field(j).Name, v.Field(j).Interface())
		}
	}

	return checkoutURL
}

// IsIPNAuthentic - check IPN authenticity (WIP)
func (checkout *YenePayCheckOut) IsIPNAuthentic(ipnModel interface{}, useSandbox bool) (bool, error) {
	ipnURL := ipnVerifyURLProd

	if useSandbox {
		ipnURL = ipnVerifyURLSandbox
	}

	reqBody, _ := json.Marshal(ipnModel)
	resp, err := http.Post(ipnURL, "application/json", bytes.NewBuffer(reqBody))

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

// RequestPDT - Request PDT model (WIP)
func (checkout *YenePayCheckOut) RequestPDT(pdtReq PdtRequest) (interface{}, error) {
	pdtURL := pdtURLProd

	if pdtReq.UseSandbox {
		pdtURL = pdtURLSandbox
	}

	reqBody, _ := json.Marshal(pdtReq)
	resp, err := http.Post(pdtURL, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		fmt.Print(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
