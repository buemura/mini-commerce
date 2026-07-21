package usecases

import (
	"context"
	"testing"

	"github.com/buemura/event-driven-commerce/mc-product-service/internal/domain/product"
	"github.com/buemura/event-driven-commerce/mc-product-service/internal/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestGetProductUsecase(t *testing.T) {
	p := []*product.Product{
		{
			ID:          1,
			Name:        "existing_product",
			Description: "existing_product_description",
			Price:       100,
			Quantity:    14,
			ImageUrl:    "http://localhost:3131",
		},
		{
			ID:          2,
			Name:        "existing_product_2",
			Description: "existing_product_description_2",
			Price:       433,
			Quantity:    122,
			ImageUrl:    "http://localhost:3131",
		},
		{
			ID:          3,
			Name:        "existing_product_3",
			Description: "existing_product_description_3",
			Price:       113,
			Quantity:    3232,
			ImageUrl:    "http://localhost:3131",
		},
	}

	repo := database.NewInMemoryProductRepo(p)
	service := NewGetProductUsecase(repo)

	t.Run("Return product struct", func(t *testing.T) {
		res, _ := service.Execute(context.Background(), 1)
		assert.Equal(t, &product.Product{
			ID:          1,
			Name:        "existing_product",
			Description: "existing_product_description",
			Price:       100,
			Quantity:    14,
			ImageUrl:    "http://localhost:3131",
		}, res)
	})
}
