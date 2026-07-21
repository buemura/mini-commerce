package contracts

import (
	"context"

	"github.com/buemura/event-driven-commerce/mc-order-service/internal/domain/product"
)

type ProductService interface {
	UpdateProductQuantity(ctx context.Context, id, quantity int) (*product.Product, error)
}
