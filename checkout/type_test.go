package checkout

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheckoutTypes(t *testing.T) {
	if !cmp.Equal(ExpressCheckout.Type, "Express") {
		t.Errorf("expected '%s' as Type, but found '%s' instead", "Express", ExpressCheckout.Type)
	}

	if !cmp.Equal(CartCheckout.Type, "Cart") {
		t.Errorf("expected '%s' as Type, but found '%s' instead", "Cart", CartCheckout.Type)
	}
}
