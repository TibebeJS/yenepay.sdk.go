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

type PdtParams PdtRequestModel

// PDT Constructor
func NewPdtRequestModel(
	params PdtParams,
) *PdtRequestModel {
	return &PdtRequestModel{
		params.PdtToken,
		params.TransactionId,
		params.MerchantId,
		params.UseSandbox,
		"PDT",
	}
}

// Validate and Marshal PDT to JSON format
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
