package checkout

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type CheckoutOption struct {
	UseSandbox            bool
	Process               CheckoutType
	MerchantId            string `validate:"required"`
	SuccessUrl            string
	CancelUrl             string
	IPNUrl                string
	FailureUrl            string
	ExpiresAfter          int
	MerchantOrderId       string
	TotalItemsDeliveryFee float64
	TotalItemsTax1        float64
	TotalItemsTax2        float64
	TotalItemsDiscount    float64
	TotalItemsHandlingFee float64
}

type OptionsParams CheckoutOption

// Return Only Required Fields for Express Checkout
// I.e. Exclude all Fields that start with 'Total'
func (self *CheckoutOption) GetExpressFields() interface{} {
	return struct {
		UseSandbox      bool
		Process         CheckoutType
		MerchantId      string
		SuccessUrl      string
		CancelUrl       string
		IPNUrl          string
		FailureUrl      string
		ExpiresAfter    int
		MerchantOrderId string
	}{
		self.UseSandbox,
		self.Process,
		self.MerchantId,
		self.SuccessUrl,
		self.CancelUrl,
		self.IPNUrl,
		self.FailureUrl,
		self.ExpiresAfter,
		self.MerchantOrderId,
	}
}

// Validate and Marshal CheckoutOption to JSON format
func (self *CheckoutOption) ToJSON(forCart bool) (string, error) {

	var fields interface{}

	if !forCart {
		fields = self.GetExpressFields()
	} else {
		fields = self
	}

	var validate *validator.Validate = validator.New()

	err := validate.Struct(self)

	if err != nil {
		return "", err
	} else {
		data, _ := json.Marshal(fields)

		return string(data), nil
	}
}

// Set Order Total Fees
// I.e. Set Fields that start with 'Total'
func (self *CheckoutOption) SetOrderFees(totalItemsDeliveryFee float64, totalItemsDiscount float64, totalItemsHandlingFee float64, totalItemsTax1, totalItemsTax2 float64) {
	self.TotalItemsDeliveryFee = totalItemsDeliveryFee
	self.TotalItemsDiscount = totalItemsDiscount
	self.TotalItemsHandlingFee = totalItemsHandlingFee
	self.TotalItemsTax1 = totalItemsTax1
	self.TotalItemsTax2 = totalItemsTax2
}

// CheckoutOption Constructor
func NewCheckoutOption(
	params OptionsParams,
) *CheckoutOption {
	return &CheckoutOption{
		params.UseSandbox,
		params.Process,
		params.MerchantId,
		params.SuccessUrl,
		params.CancelUrl,
		params.IPNUrl,
		params.FailureUrl,
		params.ExpiresAfter,
		params.MerchantOrderId,
		params.TotalItemsDeliveryFee,
		params.TotalItemsTax1,
		params.TotalItemsTax2,
		params.TotalItemsDiscount,
		params.TotalItemsHandlingFee,
	}
}
