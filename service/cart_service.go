package service

var cartItems []Menu

func AddToCart(menu Menu) {
	cartItems = append(cartItems, menu)
}

func GetCart() []Menu {
	return cartItems
}

func ClearCart() {
	cartItems = nil
}

func CartIsEmpty() bool {
	return len(cartItems) == 0
}
