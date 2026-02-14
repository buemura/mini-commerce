package contracts

import (
	"github.com/buemura/event-driven-commerce/svc-order/internal/domain/product"
)

type ProductService interface {
	UpdateProductQuantity(id, quantity int) (*product.Product, error)
}
