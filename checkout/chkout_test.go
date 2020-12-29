package checkout

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewYenePayCheckOut(t *testing.T) {
	checkout := NewYenePayCheckOut()

	expected := &YenePayCheckOut{
		"https://www.yenepay.com/checkout/Home/Process/",
		"https://test.yenepay.com/Home/Process/",
		"https://endpoints.yenepay.com/api/verify/ipn/",
		"https://testapi.yenepay.com/api/verify/ipn/",
		"https://endpoints.yenepay.com/api/verify/pdt/",
		"https://testapi.yenepay.com/api/verify/pdt/",
	}

	if !cmp.Equal(checkout, expected) {
		t.Error("NewYenePayCheckOut Constructor is not working as expected")
	}
}
