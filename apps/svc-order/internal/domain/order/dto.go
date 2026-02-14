package order

import (
	"github.com/buemura/event-driven-commerce/svc-order/internal/domain/common"
	"github.com/buemura/event-driven-commerce/svc-order/internal/domain/product"
)

type CreateOrderIn struct {
	CustomerId    string
	ProductList   []*product.Product
	PaymentMethod string
}

type GetManyOrdersIn struct {
	Page  int
	Items int
}

type UpdateOrderStatusIn struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}

type GetManyOrdersOut struct {
	OrderList []*Order
	Meta      *common.PaginationMeta
}
