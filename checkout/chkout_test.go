package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewYenePayCheckOut(t *testing.T) {
	actual := NewYenePayCheckOut()

	expected := &YenePayCheckOut{
		"https://www.yenepay.com/checkout/Home/Process/",
		"https://test.yenepay.com/Home/Process/",
		"https://endpoints.yenepay.com/api/verify/ipn/",
		"https://testapi.yenepay.com/api/verify/ipn/",
		"https://endpoints.yenepay.com/api/verify/pdt/",
		"https://testapi.yenepay.com/api/verify/pdt/",
	}

	assert.Exactly(t, expected, actual)
}

func TestGetCheckoutUrlForExpress(t *testing.T) {

	actual := NewYenePayCheckOut().GetCheckoutUrlForExpress(
		NewCheckoutOption(
			true,
			ExpressCheckout,
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
		),
		NewCheckoutItem(
			"544",
			"PC",
			30.0,
			2,
			0.2,
			0.0,
			0.0,
			0.0,
			0.0,
		),
	)

	expected := "https://test.yenepay.com/Home/Process/?CancelUrl=sdfsd&ExpiresAfter=3&FailureUrl=sdfsd&IPNUrl=sdfsd&MerchantId=222&MerchantOrderId=sdfsdsdfsd&Process=Express&SuccessUrl=sdfsd&UseSandbox=true"

	assert.Exactly(t, expected, actual)

}
