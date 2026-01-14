package controller

import (
	"net/http"
	"strconv"
	"videocall/internal/app/product/schema"
	"videocall/internal/app/product/service"
	"videocall/pkg/request"
	"videocall/pkg/response"
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

	}

	var product, err = ctrl.service.GetAllProduct()
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), product)
}

func (ctrl *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
		return
	}

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	product, err := ctrl.service.GetProductById(id)
	if err != nil {
		response.WriteResponse(w, r, http.StatusNotFound, response.CodeNotFound, response.GetMessage(response.CodeNotFound), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), product)
}

func (ctrl *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
		return
	}

	var newProductRequest *schema.ProductRequest

	err := request.DecodeRequest(r, newProductRequest)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	err = ctrl.service.CreateProduct(newProductRequest)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), newProductRequest)
}

func (ctrl *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	var newProductUpdated *schema.Product

	err = request.DecodeRequest(r, newProductUpdated)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	err = ctrl.service.UpdateProduct(id, newProductUpdated)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), newProductUpdated)
}

func (ctrl *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
	}

	err = ctrl.service.DeleteProduct(id)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
	}

	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "berhasil delete"})

}
