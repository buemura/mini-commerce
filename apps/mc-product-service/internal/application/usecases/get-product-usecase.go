package usecases

import (
	"context"

	"github.com/buemura/event-driven-commerce/mc-product-service/internal/domain/product"
)

type GetProductUsecase struct {
	repo product.ProductRepository
}

func NewGetProductUsecase(repo product.ProductRepository) *GetProductUsecase {
	return &GetProductUsecase{
		repo: repo,
	}
}

func (s *GetProductUsecase) Execute(ctx context.Context, id int) (*product.Product, error) {
	prod, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if prod == nil {
		return nil, product.ErrProductNotFound
	}
	return prod, nil
}
