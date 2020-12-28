package checkout

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheckoutTypes(t *testing.T) {
	if !cmp.Equal(string(ExpressCheckout), "Express") {
		t.Errorf("expected '%s', but found '%s' instead", "Express", ExpressCheckout)
	}

	if !cmp.Equal(string(CartCheckout), "Cart") {
		t.Errorf("expected '%s', but found '%s' instead", "Cart", CartCheckout)
	}
}
