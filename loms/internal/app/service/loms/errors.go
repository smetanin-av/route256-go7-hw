package loms

import "errors"

var (
	ErrWrongOrderStatus  = errors.New("wrong order status")
	ErrStockInsufficient = errors.New("stock insufficient")
)
