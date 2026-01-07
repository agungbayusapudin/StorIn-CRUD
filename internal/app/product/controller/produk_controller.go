package controller

import (
	"net/http"
	"strconv"
	"videocall/internal/app/product/service"
	"videocall/internal/pkg"
	"videocall/schema"
)

type productController struct {
	service service.ProductServiceInterface
}

type ProductRequest struct {
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

type ProductResponse struct {
	ID          string  `json:"id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func NewProductController(service service.ProductServiceInterface) *productController {
	return &productController{
		service: service,
	}
}

func (ctrl *productController) GetAllProduct(w http.ResponseWriter, r *http.Request) (*[]schema.Product, error) {
	if r.Method != http.MethodGet {
		pkg.WriteResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}

	var product, err = ctrl.service.GetAllProduct()
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (ctrl *productController) GetProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.WriteResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "invalid product id"})
	}

	product, err := ctrl.service.GetProductById(id)
	if err != nil {
		pkg.WriteResponse(w, http.StatusNotFound, map[string]string{"error": "product not found"})
	}

	pkg.WriteResponse(w, http.StatusOK, product)
}

func (ctrl *productController) CreateProduct(w http.ResponseWriter, r *http.Request, product *schema.Product) error {

	return nil
}

func (ctrl *productController) UpdateProduct(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (ctrl *productController) DeleteProduct(w http.ResponseWriter, r *http.Request) error {
	return nil
}
