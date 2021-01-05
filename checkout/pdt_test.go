package checkout

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPdtRequestModel(t *testing.T) {
	pdt := NewPdtRequestModel(
		PdtParams{
			PdtToken:        "test",
			TransactionId:   "1234",
			MerchantOrderId: "2345",
			UseSandbox:      true,
		},
	)

	expected := &PdtRequestModel{
		"test",
		"1234",
		"2345",
		true,
		"PDT",
	}

	assert.Exactly(t, expected, pdt)
}

func TestPDTToJSON(t *testing.T) {
	actual, _ := NewPdtRequestModel(
		PdtParams{
			PdtToken:        "test",
			TransactionId:   "1234",
			MerchantOrderId: "2345",
			UseSandbox:      true,
		},
	).ToJSON()

	expected := `
	{
		"PdtToken": "test",
		"TransactionId": "1234",
		"MerchantOrderId": "2345",
		"RequestType": "PDT"
	}
	`
	require.JSONEq(t, expected, actual)

	_, actualError := (&PdtRequestModel{}).ToJSON()

	if assert.Error(t, actualError) {
		for _, err := range actualError.(validator.ValidationErrors) {
			assert.Exactly(t, "PdtToken", err.Field())
			assert.Exactly(t, "required", err.Tag())
		}
	}
}
