package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewYenePayCheckOut(t *testing.T) {
	actual := NewYenePayCheckOut()

	expected := &YenePayCheckOut{}

	assert.Exactly(t, expected, actual)
}

func TestGetCheckoutURLForExpress(t *testing.T) {

	actual := NewYenePayCheckOut().ExpressCheckoutURL(
		NewOption(
			OptionsParams{
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
			},
		),
		NewExpressCheckoutItem(
			ExpressParams{
				ItemID:    "544",
				ItemName:  "PC",
				UnitPrice: 30.0,
				Quantity:  2,
				Discount:  0.2,
			},
		),
	)

	expected := "https://test.yenepay.com/Home/Process/?CancelURL=sdfsd&ExpiresAfter=3&FailureURL=sdfsd&IPNURL=sdfsd&MerchantID=222&MerchantOrderID=sdfsdsdfsd&Process=Express&SuccessURL=sdfsd&UseSandbox=true&DeliveryFee=0&Discount=0.2&HandlingFee=0&ItemID=544&ItemName=PC&Quantity=2&Tax1=0&Tax2=0&UnitPrice=30"

	assert.Exactly(t, expected, actual)

}
