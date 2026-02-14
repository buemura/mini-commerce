package usecases

import (
	"context"

	"github.com/buemura/event-driven-commerce/svc-product/internal/domain/product"
)

type UpdateProductQuantityUsecase struct {
	repo product.ProductRepository
}

func NewUpdateProductQuantityUsecase(repo product.ProductRepository) *UpdateProductQuantityUsecase {
	return &UpdateProductQuantityUsecase{
		repo: repo,
	}
}

func (s *UpdateProductQuantityUsecase) Execute(ctx context.Context, in *product.UpdateProductQuantityIn) (*product.Product, error) {
	prod, err := s.repo.FindById(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	if prod == nil {
		return nil, product.ErrProductNotFound
	}

	prod.Quantity += in.Quantity

	if prod.Quantity < 0 {
		return nil, product.ErrProductInsufficientQuantity
	}

	res, err := s.repo.Update(ctx, prod)
	if err != nil {
		return nil, err
	}
	return res, nil
}
