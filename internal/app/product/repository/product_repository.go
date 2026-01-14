package repository

import (
	"database/sql"
	"fmt"
	"time"
	"videocall/internal/app/product/schema"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepositoryInterface {
	return &productRepository{db: db}
}

func (repo *productRepository) GetAllProduct() ([]*schema.Product, error) {
	var result []*schema.Product

	if repo.db == nil {
		fmt.Println("Database nul")
	}

	stmt := "select * from products"
	rows, err := repo.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product schema.Product
		err := rows.Scan(&product.ID, &product.ProductName, &product.Price, &product.Stock, &product.Description, &product.CreatedAt, &product.CreatedAtUnix, &product.UpdatedAt, &product.UpdatedAtUnix)
		if err != nil {
			return nil, err
		}
		result = append(result, &product)
	}

	return result, nil
}

func (repo *productRepository) GetProductById(id int) (*schema.Product, error) {
	var product schema.Product

	stmt := "SELECT id, product_name, price FROM products WHERE id = $1"
	err := repo.db.QueryRow(stmt, id).Scan(&product.ID, &product.ProductName, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *productRepository) CreateProduct(product *schema.ProductRequest) error {
	stmt := "INSERT INTO products (product_name, price, stock, description, created_at, created_at_unix, updated_at, updated_at_unix) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := repo.db.Exec(stmt,
		product.ProductName,
		product.Price,
		product.Stock,
		product.Description,
		time.Now(),
		time.Now().UTC().UnixMilli(),
		time.Now(),
		time.Now().UTC().UnixMilli(),
	)

	fmt.Println("SELESAI REPOSITORY")
	if err != nil {
		return err
	}
	return nil
}

func (repo *productRepository) UpdateProduct(id int, product *schema.Product) error {
	stmt := "UPDATE products SET product_name = $1, price = $2 WHERE id = $3"
	_, err := repo.db.Exec(stmt, product.ProductName, product.Price, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *productRepository) DeleteProduct(id int) error {
	stmt := "DELETE FROM products WHERE id = $1"
	_, err := repo.db.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
