package usecase

import (
	productsv1 "github.com/Lineblaze/products_protos/gen/go/products"
	"golang.org/x/net/context"
	repository "products/internal"
	models2 "products/internal/models"
	"products/pkg/logger"
)

//go:generate ifacemaker -f *.go -o ../usecase.go -i UseCase -s UseCase -p internal -y "Controller describes methods, implemented by the usecase package."
type UseCase struct {
	repo   repository.Postgres
	logger *logger.ApiLogger
}

func NewUseCase(repo repository.Postgres, logger *logger.ApiLogger) *UseCase {
	return &UseCase{
		repo:   repo,
		logger: logger,
	}
}

func (u *UseCase) CreateProductCategory(ctx context.Context, input *models2.CreateProductCategoryInput) (*models2.CreateProductCategoryOutput, error) {
	category, err := u.repo.CreateProductCategory(ctx, input)
	if err != nil {
		u.logger.Errorf("Error creating product category: %v", err)
		return nil, err
	}

	return &models2.CreateProductCategoryOutput{
		Category: &productsv1.ProductCategory{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}

func (u *UseCase) GetProductCategory(ctx context.Context, id int64) (*models2.GetProductCategoryOutput, error) {
	category, err := u.repo.GetProductCategory(ctx, id)
	if err != nil {
		u.logger.Errorf("Error fetching product category: %v", err)
		return nil, err
	}

	return &models2.GetProductCategoryOutput{
		Category: &productsv1.ProductCategory{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}

func (u *UseCase) UpdateProductCategory(ctx context.Context, input *models2.UpdateProductCategoryInput) (*models2.UpdateProductCategoryOutput, error) {
	category, err := u.repo.UpdateProductCategory(ctx, input)
	if err != nil {
		u.logger.Errorf("Error updating product category: %v", err)
		return nil, err
	}

	return &models2.UpdateProductCategoryOutput{
		Category: &productsv1.ProductCategory{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}

func (u *UseCase) DeleteProductCategory(ctx context.Context, id int64) error {
	err := u.repo.DeleteProductCategory(ctx, id)
	if err != nil {
		u.logger.Errorf("Error deleting product category: %v", err)
		return err
	}
	return nil
}

func (u *UseCase) GetProductCategories(ctx context.Context) (*models2.GetProductCategoriesOutput, error) {
	categories, err := u.repo.GetProductCategories(ctx)
	if err != nil {
		u.logger.Errorf("Error fetching product categories: %v", err)
		return nil, err
	}

	var categoryList []*productsv1.ProductCategory
	for _, category := range categories {
		categoryList = append(categoryList, &productsv1.ProductCategory{
			Id:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &models2.GetProductCategoriesOutput{
		Categories: categoryList,
	}, nil
}

func (u *UseCase) CreateProduct(ctx context.Context, input *models2.CreateProductInput) (*models2.CreateProductOutput, error) {
	product, err := u.repo.CreateProduct(ctx, input)
	if err != nil {
		u.logger.Errorf("Error creating product: %v", err)
		return nil, err
	}

	return &models2.CreateProductOutput{
		Product: &productsv1.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CategoryId:  product.CategoryId,
		},
	}, nil
}

func (u *UseCase) GetProduct(ctx context.Context, id int64) (*models2.GetProductOutput, error) {
	product, err := u.repo.GetProduct(ctx, id)
	if err != nil {
		u.logger.Errorf("Error fetching product: %v", err)
		return nil, err
	}

	return &models2.GetProductOutput{
		Product: &productsv1.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CategoryId:  product.CategoryId,
		},
	}, nil
}

func (u *UseCase) UpdateProduct(ctx context.Context, input *models2.UpdateProductInput) (*models2.UpdateProductOutput, error) {
	product, err := u.repo.UpdateProduct(ctx, input)
	if err != nil {
		u.logger.Errorf("Error updating product: %v", err)
		return nil, err
	}

	return &models2.UpdateProductOutput{
		Product: &productsv1.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CategoryId:  product.CategoryId,
		},
	}, nil
}

func (u *UseCase) DeleteProduct(ctx context.Context, id int64) error {
	err := u.repo.DeleteProduct(ctx, id)
	if err != nil {
		u.logger.Errorf("Error deleting product: %v", err)
		return err
	}
	return nil
}

func (u *UseCase) GetProducts(ctx context.Context) (*models2.GetProductsOutput, error) {
	products, err := u.repo.GetProducts(ctx)
	if err != nil {
		u.logger.Errorf("Error fetching products: %v", err)
		return nil, err
	}

	var productList []*productsv1.Product
	for _, product := range products {
		productList = append(productList, &productsv1.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CategoryId:  product.CategoryId,
		})
	}

	return &models2.GetProductsOutput{
		Products: productList,
	}, nil
}
