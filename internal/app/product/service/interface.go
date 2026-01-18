package service

import "videocall/internal/app/product/schema"

type ProductServiceInterface interface {
	GetProductById(id int) (*schema.Product, error)
	GetAllProduct() ([]*schema.Product, error)
	CreateProduct(product *schema.ProductRequest) error
	UpdateProduct(id int, product *schema.Product) error
	DeleteProduct(id int) error
}
