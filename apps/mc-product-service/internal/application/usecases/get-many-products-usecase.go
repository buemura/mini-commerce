package usecases

import (
	"context"
	"math"

	"github.com/buemura/event-driven-commerce/mc-product-service/internal/domain/common"
	"github.com/buemura/event-driven-commerce/mc-product-service/internal/domain/product"
)

type GetManyProductUsecase struct {
	repo product.ProductRepository
}

func NewGetManyProductUsecase(repo product.ProductRepository) *GetManyProductUsecase {
	return &GetManyProductUsecase{
		repo: repo,
	}
}

func (uc *GetManyProductUsecase) Execute(ctx context.Context, opt *product.GetManyProductsIn) (*product.GetManyProductsOut, error) {
	res, err := uc.repo.FindMany(ctx, opt)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(res.TotalCount) / float64(opt.Items)))

	return &product.GetManyProductsOut{
		ProductList: res.ProductList,
		Meta: &common.PaginationMeta{
			Page:       opt.Page,
			Items:      opt.Items,
			TotalPages: totalPages,
			TotalItems: res.TotalCount,
		},
	}, nil
}
