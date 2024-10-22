package grpc

import (
	productsv1 "github.com/Lineblaze/products_protos/gen/go/products"
	"golang.org/x/net/context"
	"products/internal/models"
	useCase "products/internal/usecase"
	"products/pkg/logger"
)

//go:generate ifacemaker -f handler.go -o ../../handler.go -i Handler -s Handler -p internal -y "Controller describes methods, implemented by the grpc package."
type Handler struct {
	useCase *useCase.UseCase
	logger  *logger.ApiLogger
	productsv1.UnimplementedProductServiceServer
}

func NewHandler(useCase *useCase.UseCase, logger *logger.ApiLogger) *Handler {
	return &Handler{useCase: useCase, logger: logger}
}

func (h *Handler) CreateProductCategory(ctx context.Context, req *productsv1.CreateProductCategoryRequest) (*productsv1.ProductCategoryResponse, error) {
	h.logger.Infof("Creating product category: %s", req.Name)

	response, err := h.useCase.CreateProductCategory(ctx, &models.CreateProductCategoryInput{
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		h.logger.Errorf("Error creating product category: %v", err)
		return nil, err
	}

	return &productsv1.ProductCategoryResponse{
		Category: response.Category,
	}, nil
}

func (h *Handler) GetProductCategory(ctx context.Context, req *productsv1.GetProductCategoryRequest) (*productsv1.ProductCategoryResponse, error) {
	h.logger.Infof("Fetching product category with ID: %d", req.Id)

	response, err := h.useCase.GetProductCategory(ctx, req.Id)

	if err != nil {
		h.logger.Errorf("Error fetching product category: %v", err)
		return nil, err
	}

	return &productsv1.ProductCategoryResponse{
		Category: response.Category,
	}, nil
}

func (h *Handler) UpdateProductCategory(ctx context.Context, req *productsv1.UpdateProductCategoryRequest) (*productsv1.ProductCategoryResponse, error) {
	h.logger.Infof("Updating product category with ID: %d", req.Id)

	response, err := h.useCase.UpdateProductCategory(ctx, &models.UpdateProductCategoryInput{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		h.logger.Errorf("Error updating product category: %v", err)
		return nil, err
	}

	return &productsv1.ProductCategoryResponse{
		Category: response.Category,
	}, nil
}

func (h *Handler) DeleteProductCategory(ctx context.Context, req *productsv1.DeleteProductCategoryRequest) (*productsv1.DeleteProductCategoryResponse, error) {
	h.logger.Infof("Deleting product category with ID: %d", req.Id)

	err := h.useCase.DeleteProductCategory(ctx, req.Id)

	if err != nil {
		h.logger.Errorf("Error deleting product category: %v", err)
		return nil, err
	}

	return &productsv1.DeleteProductCategoryResponse{
		Message: "Product category deleted successfully.",
	}, nil
}

func (h *Handler) GetProductCategories(ctx context.Context, req *productsv1.GetProductCategoriesRequest) (*productsv1.GetProductCategoriesResponse, error) {
	h.logger.Infof("Fetching all product categories.")

	response, err := h.useCase.GetProductCategories(ctx)

	if err != nil {
		h.logger.Errorf("Error fetching product categories: %v", err)
		return nil, err
	}

	return &productsv1.GetProductCategoriesResponse{
		Categories: response.Categories,
	}, nil
}

// Product Handlers

func (h *Handler) CreateProduct(ctx context.Context, req *productsv1.CreateProductRequest) (*productsv1.ProductResponse, error) {
	h.logger.Infof("Creating product: %s", req.Name)

	response, err := h.useCase.CreateProduct(ctx, &models.CreateProductInput{
		Name:        req.Name,
		Description: req.Description,
		Price:       float64(req.Price),
		CategoryID:  req.CategoryId,
	})

	if err != nil {
		h.logger.Errorf("Error creating product: %v", err)
		return nil, err
	}

	return &productsv1.ProductResponse{
		Product: response.Product,
	}, nil
}

func (h *Handler) GetProduct(ctx context.Context, req *productsv1.GetProductRequest) (*productsv1.ProductResponse, error) {
	h.logger.Infof("Fetching product with ID: %d", req.Id)

	response, err := h.useCase.GetProduct(ctx, req.Id)

	if err != nil {
		h.logger.Errorf("Error fetching product: %v", err)
		return nil, err
	}

	return &productsv1.ProductResponse{
		Product: response.Product,
	}, nil
}

func (h *Handler) UpdateProduct(ctx context.Context, req *productsv1.UpdateProductRequest) (*productsv1.ProductResponse, error) {
	h.logger.Infof("Updating product with ID: %d", req.Id)

	response, err := h.useCase.UpdateProduct(ctx, &models.UpdateProductInput{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       float64(req.Price),
		CategoryID:  req.CategoryId,
	})

	if err != nil {
		h.logger.Errorf("Error updating product: %v", err)
		return nil, err
	}

	return &productsv1.ProductResponse{
		Product: response.Product,
	}, nil
}

func (h *Handler) DeleteProduct(ctx context.Context, req *productsv1.DeleteProductRequest) (*productsv1.DeleteProductResponse, error) {
	h.logger.Infof("Deleting product with ID: %d", req.Id)

	err := h.useCase.DeleteProduct(ctx, req.Id)

	if err != nil {
		h.logger.Errorf("Error deleting product: %v", err)
		return nil, err
	}

	return &productsv1.DeleteProductResponse{
		Message: "Product deleted successfully.",
	}, nil
}

func (h *Handler) GetProducts(ctx context.Context, req *productsv1.GetProductsRequest) (*productsv1.GetProductsResponse, error) {
	h.logger.Infof("Fetching all products.")

	response, err := h.useCase.GetProducts(ctx)

	if err != nil {
		h.logger.Errorf("Error fetching products: %v", err)
		return nil, err
	}

	return &productsv1.GetProductsResponse{
		Products: response.Products,
	}, nil
}
