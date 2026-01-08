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
	mux.HandleFunc("GET /", r.controller.GetAllProduct)
	mux.HandleFunc("GET /{id}", r.controller.GetProductById)
	mux.HandleFunc("POST /", r.controller.CreateProduct)
	mux.HandleFunc("PUT /{id}", r.controller.UpdateProduct)
	mux.HandleFunc("DELETE /{id}", r.controller.DeleteProduct)
}
