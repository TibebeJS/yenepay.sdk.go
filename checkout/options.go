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

// Return Only Required Fields for Cart Checkout
// I.e. Exclude all Fields that start with 'Total'
func (self *CheckoutOption) GetCartFields() interface{} {
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

	if forCart {
		fields = self.GetCartFields()
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
	UseSandbox bool,
	Process CheckoutType,
	MerchantId string,
	SuccessUrl string,
	CancelUrl string,
	IPNUrl string,
	FailureUrl string,
	ExpiresAfter int,
	MerchantOrderId string,
	TotalItemsDeliveryFee float64,
	TotalItemsTax1 float64,
	TotalItemsTax2 float64,
	TotalItemsDiscount float64,
	TotalItemsHandlingFee float64,
) *CheckoutOption {
	return &CheckoutOption{
		UseSandbox,
		Process,
		MerchantId,
		SuccessUrl,
		CancelUrl,
		IPNUrl,
		FailureUrl,
		ExpiresAfter,
		MerchantOrderId,
		TotalItemsDeliveryFee,
		TotalItemsTax1,
		TotalItemsTax2,
		TotalItemsDiscount,
		TotalItemsHandlingFee,
	}
}
