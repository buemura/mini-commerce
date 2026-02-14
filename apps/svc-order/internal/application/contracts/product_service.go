package contracts

import (
	"context"

	"github.com/buemura/event-driven-commerce/svc-order/internal/domain/product"
)

type ProductService interface {
	UpdateProductQuantity(ctx context.Context, id, quantity int) (*product.Product, error)
}
