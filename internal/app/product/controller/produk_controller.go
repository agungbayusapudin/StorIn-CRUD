package controller

import (
	"net/http"
	"strconv"
	"videocall/internal/app/product/schema"
	"videocall/internal/app/product/service"
	"videocall/pkg"
)

type ProductController struct {
	service service.ProductServiceInterface
}

func NewProductController(service service.ProductServiceInterface) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (ctrl *ProductController) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.WriteResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}

	var product, err = ctrl.service.GetAllProduct()
	if err != nil {
		pkg.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	pkg.WriteResponse(w, http.StatusOK, product)
}

func (ctrl *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
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

func (ctrl *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		pkg.WriteResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}

	var newProductRequest *schema.ProductRequest

	err := pkg.DecodeResposne(r, newProductRequest)
	if err != nil {
		pkg.WriteResponse(w, http.StatusBadRequest, newProductRequest)
	}

	err = ctrl.service.CreateProduct(newProductRequest)
	if err != nil {
		pkg.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "nternal Server Error"})
	}

	pkg.WriteResponse(w, http.StatusCreated, map[string]string{"status": "berhasil created"})
}

func (ctrl *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		pkg.WriteResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method Not Allowed"})
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.WriteResponse(w, http.StatusNotFound, map[string]string{"error": "Url Path salah"})
	}

	var newProductUpdated *schema.Product

	err = pkg.DecodeResposne(r, newProductUpdated)
	if err != nil {
		pkg.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Input Salah"})
	}

	err = ctrl.service.UpdateProduct(id, newProductUpdated)
	if err != nil {
		pkg.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	pkg.WriteResponse(w, http.StatusOK, newProductUpdated)
}

func (ctrl *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		pkg.WriteResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method Not Allowed"})
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkg.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Url Path salah"})
	}

	err = ctrl.service.DeleteProduct(id)
	if err != nil {
		pkg.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Status Internal Server Eror"})
	}

	pkg.WriteResponse(w, http.StatusOK, map[string]string{"status": "Berhasil Detele Product"})

}
