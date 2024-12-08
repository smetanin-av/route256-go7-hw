package loms

import (
	"log"

	"github.com/samber/lo"
	"route256/loms/internal/app/domain"
	"route256/loms/pkg/loms/v1"
)

func fromCreateOrderRequest(req *loms_v1.CreateOrderRequest) []domain.OrderItem {
	return lo.Map(req.Items, fromCreateOrderRequestItem)
}

func fromCreateOrderRequestItem(item *loms_v1.CreateOrderRequestItem, _ int) domain.OrderItem {
	return domain.OrderItem{
		SKU:   item.Sku,
		Count: uint16(item.Count),
	}
}

func toOrderStatus(status domain.Status) loms_v1.OrderStatus {
	switch status {
	case domain.StatusNew:
		return loms_v1.OrderStatus_NEW

	case domain.StatusAwaitingPayment:
		return loms_v1.OrderStatus_AWAITING_PAYMENT

	case domain.StatusFailed:
		return loms_v1.OrderStatus_FAILED

	case domain.StatusPayed:
		return loms_v1.OrderStatus_PAYED

	case domain.StatusCancelled:
		return loms_v1.OrderStatus_CANCELLED

	default:
		log.Printf("WARN invalid status %d", status)
		return loms_v1.OrderStatus_INVALID
	}
}

func toListOrderResponse(mdl *domain.OrderInfo) *loms_v1.ListOrderResponse {
	return &loms_v1.ListOrderResponse{
		Status: toOrderStatus(mdl.Status),
		User:   mdl.UserID,
		Items:  lo.Map(mdl.Items, toListOrderResponseItem),
	}
}

func toListOrderResponseItem(item *domain.OrderItem, _ int) *loms_v1.ListOrderResponseItem {
	return &loms_v1.ListOrderResponseItem{
		Sku:   item.SKU,
		Count: uint32(item.Count),
	}
}

func toGetStocksResponse(stocks []*domain.StockInfo) *loms_v1.GetStocksResponse {
	return &loms_v1.GetStocksResponse{
		Stocks: lo.Map(stocks, toGetStocksResponseItem),
	}
}

func toGetStocksResponseItem(stock *domain.StockInfo, _ int) *loms_v1.GetStocksResponseItem {
	return &loms_v1.GetStocksResponseItem{
		WarehouseId: stock.WarehouseID,
		Count:       stock.Count,
	}
}
