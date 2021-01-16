package checkout

// Option - checkout option model
type Option struct {
	UseSandbox            bool
	Process               Type
	MerchantID            string `validate:"required"`
	SuccessURL            string
	CancelURL             string
	IPNURL                string
	FailureURL            string
	ExpiresAfter          int
	MerchantOrderID       string
	TotalItemsDeliveryFee float64
	TotalItemsTax1        float64
	TotalItemsTax2        float64
	TotalItemsDiscount    float64
	TotalItemsHandlingFee float64
}

// OptionsParams - arugment for constructor
type OptionsParams Option

// GetExpressFields - Return Only Required Fields for Express Checkout
// I.e. Exclude all Fields that start with 'Total'
func (option *Option) GetExpressFields() interface{} {
	return struct {
		UseSandbox      bool
		Process         Type
		MerchantID      string
		SuccessURL      string
		CancelURL       string
		IPNURL          string
		FailureURL      string
		ExpiresAfter    int
		MerchantOrderID string
	}{
		option.UseSandbox,
		option.Process,
		option.MerchantID,
		option.SuccessURL,
		option.CancelURL,
		option.IPNURL,
		option.FailureURL,
		option.ExpiresAfter,
		option.MerchantOrderID,
	}
}

// // Validate and Marshal Option to JSON format
// func (option *Option) ToJSON(forCart bool) (string, error) {

// 	var fields interface{}

// 	if !forCart {
// 		fields = option.GetExpressFields()
// 	} else {
// 		fields = option
// 	}

// 	var validate *validator.Validate = validator.New()

// 	err := validate.Struct(option)

// 	if err != nil {
// 		return "", err
// 	} else {
// 		data, _ := json.Marshal(fields)

// 		return string(data), nil
// 	}
// }

// SetOrderFees - Set Order Total Fees
// I.e. Set Fields that start with 'Total'
func (option *Option) SetOrderFees(totalItemsDeliveryFee float64, totalItemsDiscount float64, totalItemsHandlingFee float64, totalItemsTax1, totalItemsTax2 float64) {
	option.TotalItemsDeliveryFee = totalItemsDeliveryFee
	option.TotalItemsDiscount = totalItemsDiscount
	option.TotalItemsHandlingFee = totalItemsHandlingFee
	option.TotalItemsTax1 = totalItemsTax1
	option.TotalItemsTax2 = totalItemsTax2
}

// NewOption - Option Constructor
func NewOption(
	params OptionsParams,
) *Option {
	return &Option{
		params.UseSandbox,
		params.Process,
		params.MerchantID,
		params.SuccessURL,
		params.CancelURL,
		params.IPNURL,
		params.FailureURL,
		params.ExpiresAfter,
		params.MerchantOrderID,
		params.TotalItemsDeliveryFee,
		params.TotalItemsTax1,
		params.TotalItemsTax2,
		params.TotalItemsDiscount,
		params.TotalItemsHandlingFee,
	}
}
