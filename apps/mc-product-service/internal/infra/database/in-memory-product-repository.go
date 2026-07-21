package database

import (
	"context"

	"github.com/buemura/event-driven-commerce/mc-product-service/internal/domain/product"
)

type InMemoryProductRepo struct {
	prod []*product.Product
}

func NewInMemoryProductRepo(prod []*product.Product) *InMemoryProductRepo {
	return &InMemoryProductRepo{
		prod: prod,
	}
}

func (r *InMemoryProductRepo) FindMany(_ context.Context, in *product.GetManyProductsIn) (*product.ProductRepositoryPaginatedOut, error) {
	return &product.ProductRepositoryPaginatedOut{
		ProductList: r.prod,
		TotalCount:  len(r.prod),
	}, nil
}

func (r *InMemoryProductRepo) FindById(_ context.Context, id int) (*product.Product, error) {
	var p *product.Product
	for _, v := range r.prod {
		if v.ID == id {
			p = v
			break
		}
	}
	return p, nil
}

func (r *InMemoryProductRepo) Update(_ context.Context, newP *product.Product) (*product.Product, error) {
	var p *product.Product
	for _, v := range r.prod {
		if v.ID == newP.ID {
			p = v
			v = newP
			break
		}
	}

	if p == nil {
		r.prod = append(r.prod, p)
	}
	return newP, nil
}
