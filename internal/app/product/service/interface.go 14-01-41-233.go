package service

import "videocall/schema"

type ProductServiceInterface interface {
	GetProductById(id int) (schema.Product, error)
	GetAllProduct() ([]schema.Product, error)
	CreateProduct(product schema.Product) error
	UpdateProduct(id int, product schema.Product) error
	DeleteProduct(id int) error
}
