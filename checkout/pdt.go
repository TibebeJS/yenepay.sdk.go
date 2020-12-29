package checkout

import "encoding/json"

type PdtRequestModel struct {
	PdtToken      string
	TransactionId string
	MerchantId    string
	UseSandbox    bool `json:"-"`
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

func (self *PdtRequestModel) ToJSON() (string, error) {
	data, err := json.Marshal(self)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
