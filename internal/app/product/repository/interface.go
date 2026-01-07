package repository

import "videocall/schema"

type ProductRepositoryInterface interface {
	GetProductById(id int) (schema.Product, error)
	GetAllProduct() ([]schema.Product, error)
	CreateProduct(product schema.Product) error
	UpdateProduct(id int, product schema.Product) error
	DeleteProduct(id int) error
}
