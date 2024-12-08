package checkout

import "errors"

var (
	ErrStockInsufficient = errors.New("stock insufficient")
	ErrCartIsEmpty       = errors.New("cart is empty")
	ErrCartInsufficient  = errors.New("cart insufficient")
)
