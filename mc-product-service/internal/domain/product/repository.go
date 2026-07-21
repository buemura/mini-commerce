package product

import "context"

type ProductRepositoryPaginatedOut struct {
	ProductList []*Product
	TotalCount  int
}

type ProductRepository interface {
	FindMany(ctx context.Context, in *GetManyProductsIn) (*ProductRepositoryPaginatedOut, error)
	FindById(ctx context.Context, id int) (*Product, error)
	Update(ctx context.Context, p *Product) (*Product, error)
}
