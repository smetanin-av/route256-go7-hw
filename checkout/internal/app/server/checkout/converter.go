package checkout

import (
	"github.com/samber/lo"
	"route256/checkout/internal/app/domain"
	"route256/checkout/pkg/checkout/v1"
)

func fromAddToCartRequest(req *checkout_v1.AddToCartRequest) *domain.UpdateCart {
	return &domain.UpdateCart{
		UserID: req.User,
		SKU:    req.Sku,
		Count:  uint16(req.Count),
	}
}

func fromDeleteFromCartRequest(req *checkout_v1.DeleteFromCartRequest) *domain.UpdateCart {
	return &domain.UpdateCart{
		UserID: req.User,
		SKU:    req.Sku,
		Count:  uint16(req.Count),
	}
}

func toListCartResponse(mdl *domain.ListCart) *checkout_v1.ListCartResponse {
	return &checkout_v1.ListCartResponse{
		Items:      lo.Map(mdl.Items, toListCartResponseItem),
		TotalPrice: mdl.TotalPrice,
	}
}

func toListCartResponseItem(item *domain.ListCartItem, _ int) *checkout_v1.ListCartResponseItem {
	return &checkout_v1.ListCartResponseItem{
		Sku:   item.SKU,
		Count: uint32(item.Count),
		Name:  item.Name,
		Price: item.Price,
	}
}
