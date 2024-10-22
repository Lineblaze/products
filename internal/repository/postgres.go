package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	productsv1 "github.com/Lineblaze/products_protos/gen/go/products"
	"golang.org/x/net/context"
	"products/internal/models"
	"products/pkg/logger"
	"products/pkg/storage/postgres"
)

//go:generate ifacemaker -f postgres.go -o ../postgres.go -i Postgres -s Postgres -p internal -y "Controller describes methods, implemented by the Postgres package."
type Postgres struct {
	db     postgres.Postgres
	logger *logger.ApiLogger
}

func NewPostgresRepository(db postgres.Postgres, logger *logger.ApiLogger) *Postgres {
	return &Postgres{db: db, logger: logger}
}

func (r *Postgres) CreateProductCategory(ctx context.Context, input *models.CreateProductCategoryInput) (*productsv1.ProductCategory, error) {
	var category productsv1.ProductCategory

	query := `INSERT INTO product_categories (name, description) VALUES ($1, $2) RETURNING id, name, description`
	err := r.db.QueryRowContext(ctx, query, input.Name, input.Description).Scan(&category.Id, &category.Name, &category.Description)
	if err != nil {
		r.logger.Errorf("Error creating product category: %v", err)
		return nil, err
	}

	return &category, nil
}

func (r *Postgres) GetProductCategory(ctx context.Context, id int64) (*productsv1.ProductCategory, error) {
	var category productsv1.ProductCategory

	query := `SELECT id, name, description FROM product_categories WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&category.Id, &category.Name, &category.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("category not found")
		}
		r.logger.Errorf("Error fetching product category: %v", err)
		return nil, err
	}

	return &category, nil
}

func (r *Postgres) UpdateProductCategory(ctx context.Context, input *models.UpdateProductCategoryInput) (*productsv1.ProductCategory, error) {
	var category productsv1.ProductCategory

	query := `UPDATE product_categories SET name = $1, description = $2 WHERE id = $3 RETURNING id, name, description`
	err := r.db.QueryRowContext(ctx, query, input.Name, input.Description, input.ID).Scan(&category.Id, &category.Name, &category.Description)
	if err != nil {
		r.logger.Errorf("Error updating product category: %v", err)
		return nil, err
	}

	return &category, nil
}

func (r *Postgres) DeleteProductCategory(ctx context.Context, id int64) error {
	query := `DELETE FROM product_categories WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Errorf("Error deleting product category: %v", err)
		return err
	}
	return nil
}

func (r *Postgres) GetProductCategories(ctx context.Context) ([]*productsv1.ProductCategory, error) {
	var categories []*productsv1.ProductCategory

	query := `SELECT id, name, description FROM product_categories`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		r.logger.Errorf("Error fetching product categories: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category productsv1.ProductCategory
		if err = rows.Scan(&category.Id, &category.Name, &category.Description); err != nil {
			r.logger.Errorf("Error scanning product category row: %v", err)
			return nil, err
		}
		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		r.logger.Errorf("Error in rows iteration: %v", err)
		return nil, err
	}

	return categories, nil
}

func (r *Postgres) CreateProduct(ctx context.Context, input *models.CreateProductInput) (*productsv1.Product, error) {
	var product productsv1.Product

	query := `INSERT INTO products (name, description, price, category_id) VALUES ($1, $2, $3, $4) RETURNING id, name, description, price, category_id`
	err := r.db.QueryRowContext(ctx, query, input.Name, input.Description, input.Price, input.CategoryID).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId)
	if err != nil {
		r.logger.Errorf("Error creating product: %v", err)
		return nil, err
	}

	return &product, nil
}

func (r *Postgres) GetProduct(ctx context.Context, id int64) (*productsv1.Product, error) {
	var product productsv1.Product

	query := `SELECT id, name, description, price, category_id FROM products WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product not found")
		}
		r.logger.Errorf("Error fetching product: %v", err)
		return nil, err
	}

	return &product, nil
}

func (r *Postgres) UpdateProduct(ctx context.Context, input *models.UpdateProductInput) (*productsv1.Product, error) {
	var product productsv1.Product

	query := `UPDATE products SET name = $1, description = $2, price = $3, category_id = $4 WHERE id = $5 RETURNING id, name, description, price, category_id`
	err := r.db.QueryRowContext(ctx, query, input.Name, input.Description, input.Price, input.CategoryID, input.ID).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId)
	if err != nil {
		r.logger.Errorf("Error updating product: %v", err)
		return nil, err
	}

	return &product, nil
}

func (r *Postgres) DeleteProduct(ctx context.Context, id int64) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Errorf("Error deleting product: %v", err)
		return err
	}
	return nil
}

func (r *Postgres) GetProducts(ctx context.Context) ([]*productsv1.Product, error) {
	var products []*productsv1.Product

	query := `SELECT id, name, description, price, category_id FROM products`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		r.logger.Errorf("Error fetching products: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product productsv1.Product
		if err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId); err != nil {
			r.logger.Errorf("Error scanning product row: %v", err)
			return nil, err
		}
		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		r.logger.Errorf("Error in rows iteration: %v", err)
		return nil, err
	}

	return products, nil
}
