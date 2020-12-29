package checkout

import "encoding/json"

type PdtRequestModel struct {
	PdtToken      string
	TransactionId string
	MerchantId    string
	UseSandbox    bool
	RequestType   string
}

func NewPdtRequestModel(
	PdtToken string,
	TransactionId string,
	MerchantId string,
	UseSandbox bool,
) *PdtRequestModel {
	return &PdtRequestModel{
		PdtToken,
		TransactionId,
		MerchantId,
		UseSandbox,
		"PDT",
	}
}

func (self *PdtRequestModel) ToJSON() ([]byte, error) {
	return json.Marshal(struct {
		PdtToken      string
		TransactionId string
		MerchantId    string
		RequestType   string
	}{
		self.PdtToken,
		self.TransactionId,
		self.MerchantId,
		self.RequestType,
	})
}
