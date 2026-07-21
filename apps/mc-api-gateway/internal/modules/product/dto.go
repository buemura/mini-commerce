package product

import (
	"github.com/buemura/event-driven-commerce/mc-api-gateway/internal/shared"
)

type GetManyProductsIn struct {
	Page  int `json:"page"`
	Items int `json:"items"`
}

type GetManyProductsOut struct {
	ProductList []*Product             `json:"product_list"`
	Meta        *shared.PaginationMeta `json:"meta"`
}
