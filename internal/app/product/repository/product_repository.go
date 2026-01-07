package repository

import (
	"database/sql"
	"videocall/schema"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (repo *productRepository) GetAllProducts() (*[]schema.Product, error) {
	var result []schema.Product

	stmt := "SELECT id, name, price FROM products"
	rows, err := repo.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product schema.Product
		err := rows.Scan(&product.ID, &product.ProductName, &product.Price)
		if err != nil {
			return nil, err
		}
		result = append(result, product)
	}

	return &result, nil
}

func (repo *productRepository) GetProductById(id int) (*schema.Product, error) {
	var product schema.Product

	stmt := "SELECT id, name, price FROM products WHERE id = $1"
	err := repo.db.QueryRow(stmt, id).Scan(&product.ID, &product.ProductName, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *productRepository) CreateProduct(product schema.Product) error {
	stmt := "INSERT INTO products (name, price) VALUES ($1, $2)"
	_, err := repo.db.Exec(stmt, product.ProductName, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (repo *productRepository) UpdateProduct(id int, product schema.Product) error {
	stmt := "UPDATE products SET name = $1, price = $2 WHERE id = $3"
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
