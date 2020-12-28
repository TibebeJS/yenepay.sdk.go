package checkout

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheckoutTypes(t *testing.T) {
	if !cmp.Equal(fmt.Sprintf("%s", ExpressCheckout), "Express") {
		t.Errorf("expected '%s', but found '%s' instead", "Express", ExpressCheckout)
	}

	if !cmp.Equal(fmt.Sprintf("%s", CartCheckout), "Cart") {
		t.Errorf("expected '%s', but found '%s' instead", "Cart", CartCheckout)
	}
}
