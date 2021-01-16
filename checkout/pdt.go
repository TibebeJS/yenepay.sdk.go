package checkout

// {
// 	RequestType: 'PDT',
// 	PdtToken: 'VEzzj2vH8mKPde',
// 	TransactionID: '290f0ecf-ee7d-4684-9841-5c0e0687920a',
// 	MerchantOrderID: '12-34',
// 	UseSandbox: true,
// 	GetPDTDictionary: [Function (anonymous)]
// }

// PdtRequest - request model for transaction verification
type PdtRequest struct {
	PdtToken        string `validate:"required,min=1"`
	TransactionID   string
	MerchantOrderID string
	UseSandbox      bool `json:"-"`
	RequestType     string
}

// {
// 	result: 'SUCCESS',
// 	TotalAmount: '18.50',
// 	BuyerID: 'f60502c5-12c8-4f2a-8709-9679fe4a386f',
// 	MerchantOrderID: '12-34',
// 	MerchantCode: '0694',
// 	MerchantID: '333dae30-2611-4fe9-aa7c-d58ea684f041',
// 	TransactionCode: 'NSLFTQLF',
// 	TransactionID: '290f0ecf-ee7d-4684-9841-5c0e0687920a',
// 	Status: 'Paid',
// 	Currency: 'ETB'
// }

// PdtResponse - result of transaction verificaiton(pdt request)
type PdtResponse struct {
	PdtToken        string `validate:"required,min=1"`
	TransactionID   string
	MerchantOrderID string
	UseSandbox      bool `json:"-"`
	RequestType     string
}

// PdtParams - arguments for PDT request constructor.
type PdtParams PdtRequest

// NewPdtRequest - PDT Constructor
func NewPdtRequest(
	params PdtParams,
) *PdtRequest {
	return &PdtRequest{
		params.PdtToken,
		params.TransactionID,
		params.MerchantOrderID,
		params.UseSandbox,
		"PDT",
	}
}
