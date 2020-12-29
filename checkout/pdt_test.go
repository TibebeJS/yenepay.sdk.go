package checkout

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPdtRequestModel(t *testing.T) {
	pdt := NewPdtRequestModel(
		"test",
		"1234",
		"2345",
		true,
	)

	expected := &PdtRequestModel{
		"test",
		"1234",
		"2345",
		true,
		"PDT",
	}

	assert.Exactly(t, pdt, expected)
}

func TestPDTToJSON(t *testing.T) {
	actual, _ := NewPdtRequestModel(
		"test",
		"1234",
		"2345",
		true,
	).ToJSON()

	expected := `
	{
		"PdtToken": "test",
		"TransactionId": "1234",
		"MerchantId": "2345",
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
