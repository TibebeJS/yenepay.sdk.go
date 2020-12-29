package checkout

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type PdtRequestModel struct {
	PdtToken      string `validate:"required,min=1"`
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
	var validate *validator.Validate = validator.New()

	err := validate.Struct(self)

	if err != nil {
		return "", err
	} else {
		data, _ := json.Marshal(self)

		return string(data), nil
	}

}
