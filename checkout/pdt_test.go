package checkout

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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

	if !cmp.Equal(pdt, expected) {
		t.Error("NewPdtRequestModel Constructor is not working as expected")
	}
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

}
