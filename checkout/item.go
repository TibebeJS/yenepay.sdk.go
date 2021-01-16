package checkout

// ExpressCheckoutItem - model
type ExpressCheckoutItem struct {
	ItemID      string  `validate:"required,min=1"`
	ItemName    string  `validate:"required,min=1"`
	UnitPrice   float64 `validate:"min=0.1"`
	Quantity    int     `validate:"min=1"`
	Discount    float64 `validate:"min=0.0"`
	HandlingFee float64 `validate:"min=0.0"`
	DeliveryFee float64 `validate:"min=0.0"`
	Tax1        float64 `validate:"min=0.0"`
	Tax2        float64 `validate:"min=0.0"`
}

// // Validate and Marshal CheckoutItem to JSON format
// func (self *ExpressCheckoutItem) ToJSON() (string, error) {
// 	var validate *validator.Validate = validator.New()

// 	err := validate.Struct(self)

// 	if err != nil {
// 		return "", err
// 	} else {
// 		data, _ := json.Marshal(self)

// 		return string(data), nil
// 	}
// }

// ExpressParams - argument type for constructor
type ExpressParams ExpressCheckoutItem

// NewExpressCheckoutItem - ExpressCheckoutItem Constructor
func NewExpressCheckoutItem(
	params ExpressParams,
) *ExpressCheckoutItem {
	return &ExpressCheckoutItem{
		params.ItemID,
		params.ItemName,
		params.UnitPrice,
		params.Quantity,
		params.Discount,
		params.HandlingFee,
		params.DeliveryFee,
		params.Tax1,
		params.Tax2,
	}
}

// CartCheckoutItem - model
type CartCheckoutItem struct {
	ItemID    string
	ItemName  string
	UnitPrice float64
	Quantity  int
}

// CartParams - argument for constructor
type CartParams CartCheckoutItem

// NewCartCheckoutItem - CartCheckoutItem Constructor
func NewCartCheckoutItem(params CartParams) *CartCheckoutItem {
	return &CartCheckoutItem{
		params.ItemID,
		params.ItemName,
		params.UnitPrice,
		params.Quantity,
	}
}
