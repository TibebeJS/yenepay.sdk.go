package checkout

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
