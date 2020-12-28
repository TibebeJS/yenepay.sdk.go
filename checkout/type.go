package checkout

type CheckoutType struct {
	Type string
}

var ExpressCheckout = CheckoutType{
	Type: "Express",
}

var CartCheckout = CheckoutType{
	Type: "Cart",
}
