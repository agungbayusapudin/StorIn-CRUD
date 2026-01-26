package routers

import (
	"net/http"
	"videocall/internal/app/product/controller"
)

type ProductRouter struct {
	controller *controller.ProductController
}

func NewProductRouters(ctrl *controller.ProductController) *ProductRouter {
	return &ProductRouter{controller: ctrl}
}

func (r *ProductRouter) RegisterRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /product", r.controller.GetAllProduct)
	mux.HandleFunc("GET /product/{id}", r.controller.GetProductById)
	mux.HandleFunc("POST /product", r.controller.CreateProduct)
	mux.HandleFunc("PUT /product/{id}", r.controller.UpdateProduct)
	mux.HandleFunc("DELETE /product/{id}", r.controller.DeleteProduct)
}
