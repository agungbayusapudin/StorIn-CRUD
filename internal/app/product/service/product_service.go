package service

import (
	"videocall/internal/app/product/repository"
	"videocall/internal/app/product/schema"
)

type productService struct {
	repo repository.ProductRepositoryInterface
}

func NewProductService(repo repository.ProductRepositoryInterface) ProductServiceInterface {
	return &productService{repo: repo}
}

func (ser *productService) GetAllProduct() ([]*schema.Product, error) {
	var product, err = ser.repo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ser *productService) GetProductById(id int) (*schema.Product, error) {
	var product, err = ser.repo.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ser *productService) CreateProduct(product *schema.ProductRequest) error {
	var err = ser.repo.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (ser *productService) UpdateProduct(id int, product *schema.Product) error {
	var err = ser.repo.UpdateProduct(id, product)
	if err != nil {
		return err
	}

	return nil
}

func (ser *productService) DeleteProduct(id int) error {
	var err = ser.repo.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
