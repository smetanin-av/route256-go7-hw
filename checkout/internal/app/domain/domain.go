package domain

type UpdateCart struct {
	UserID int64
	SKU    uint32
	Count  uint16
}

type StockInfo struct {
	WarehouseID int64
	Count       uint64
}

type ProductInfo struct {
	Name  string
	Price uint32
}

type OrderItem struct {
	SKU   uint32
	Count uint16
}

type OrderInfo struct {
	UserID int64
	Items  []*OrderItem
}

type ListCartItem struct {
	SKU   uint32
	Count uint16
	Name  string
	Price uint32
}

type ListCart struct {
	Items      []*ListCartItem
	TotalPrice uint32
}
