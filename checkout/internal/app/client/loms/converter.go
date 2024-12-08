package loms

import (
	"route256/checkout/internal/app/domain"
	"route256/checkout/internal/pb/route256/loms/v1"
)

func toCreateOrderRequest(mdl *domain.OrderInfo) *loms_v1.CreateOrderRequest {
	req := loms_v1.CreateOrderRequest{
		User:  mdl.UserID,
		Items: make([]*loms_v1.CreateOrderRequestItem, 0, len(mdl.Items)),
	}

	for _, item := range mdl.Items {
		req.Items = append(req.Items, toOrderItem(item))
	}

	return &req
}

func toOrderItem(item *domain.OrderItem) *loms_v1.CreateOrderRequestItem {
	return &loms_v1.CreateOrderRequestItem{
		Sku:   item.SKU,
		Count: uint32(item.Count),
	}
}

func fromGetStocksResponse(res *loms_v1.GetStocksResponse) []domain.StockInfo {
	stocks := make([]domain.StockInfo, 0, len(res.Stocks))
	for _, stock := range res.Stocks {
		stocks = append(stocks, domain.StockInfo{
			WarehouseID: stock.WarehouseId,
			Count:       stock.Count,
		})
	}

	return stocks
}
