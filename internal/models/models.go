package models

import productsv1 "github.com/Lineblaze/products_protos/gen/go/products"

type CreateProductCategoryInput struct {
	Name        string
	Description string
}

type CreateProductCategoryOutput struct {
	Category *productsv1.ProductCategory
}

type GetProductCategoryOutput struct {
	Category *productsv1.ProductCategory
}

type UpdateProductCategoryInput struct {
	ID          int64
	Name        string
	Description string
}

type UpdateProductCategoryOutput struct {
	Category *productsv1.ProductCategory
}

type GetProductCategoriesOutput struct {
	Categories []*productsv1.ProductCategory
}

type CreateProductInput struct {
	Name        string
	Description string
	Price       float64
	CategoryID  int64
}

type CreateProductOutput struct {
	Product *productsv1.Product
}

type GetProductOutput struct {
	Product *productsv1.Product
}

type UpdateProductInput struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	CategoryID  int64
}

type UpdateProductOutput struct {
	Product *productsv1.Product
}

type GetProductsOutput struct {
	Products []*productsv1.Product
}
