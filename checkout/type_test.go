package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckoutTypes(t *testing.T) {
	assert.Exactly(t, "Express", string(ExpressCheckout))
	assert.Exactly(t, "Cart", string(CartCheckout))
}
