package checkout

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheckoutItemConstructor(t *testing.T) {
	item := NewCheckoutItem(
		"2",
		"item-name",
		10.0,
		2,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	)

	expected := &CheckoutItem{
		"2",
		"item-name",
		10.0,
		2,
		0.0,
		0.0,
		0.0,
		0.0,
		0.0,
	}

	if !cmp.Equal(item, expected) {
		t.Error("Constructor is not working as expected")
	}
}
