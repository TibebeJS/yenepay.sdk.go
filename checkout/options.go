package checkout

import "encoding/json"

type CheckoutOption struct {
	UseSandbox            bool
	Process               CheckoutType
	MerchantId            string
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

func (self *CheckoutOption) ToJSON(forCart bool) (string, error) {

	var fields interface{}

	if forCart {
		fields = self.GetCartFields()
	} else {
		fields = self
	}

	data, err := json.Marshal(fields)

	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (self *CheckoutOption) SetOrderFees(totalItemsDeliveryFee float64, totalItemsDiscount float64, totalItemsHandlingFee float64, totalItemsTax1, totalItemsTax2 float64) {
	self.TotalItemsDeliveryFee = totalItemsDeliveryFee
	self.TotalItemsDiscount = totalItemsDiscount
	self.TotalItemsHandlingFee = totalItemsHandlingFee
	self.TotalItemsTax1 = totalItemsTax1
	self.TotalItemsTax2 = totalItemsTax2
}

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
