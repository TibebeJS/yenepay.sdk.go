package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPdtRequest(t *testing.T) {
	pdt := NewPdtRequest(
		PdtParams{
			PdtToken:        "test",
			TransactionID:   "1234",
			MerchantOrderID: "2345",
			UseSandbox:      true,
		},
	)

	expected := &PdtRequest{
		"test",
		"1234",
		"2345",
		true,
		"PDT",
	}

	assert.Exactly(t, expected, pdt)
}

// func TestPDTToJSON(t *testing.T) {
// 	actual, _ := NewPdtRequest(
// 		PdtParams{
// 			PdtToken:        "test",
// 			TransactionID:   "1234",
// 			MerchantOrderID: "2345",
// 			UseSandbox:      true,
// 		},
// 	).ToJSON()

// 	expected := `
// 	{
// 		"PdtToken": "test",
// 		"TransactionID": "1234",
// 		"MerchantOrderID": "2345",
// 		"RequestType": "PDT"
// 	}
// 	`
// 	require.JSONEq(t, expected, actual)

// 	_, actualError := (&PdtRequestModel{}).ToJSON()

// 	if assert.Error(t, actualError) {
// 		for _, err := range actualError.(validator.ValidationErrors) {
// 			assert.Exactly(t, "PdtToken", err.Field())
// 			assert.Exactly(t, "required", err.Tag())
// 		}
// 	}
// }
