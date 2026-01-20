package routers

import (
	"net/http"
	"videocall/internal/app/orders/controller"
)

type OrderRouter struct {
	controller *controller.OrderController
}

func NewOrderRouters(controller *controller.OrderController) *OrderRouter {
	return &OrderRouter{controller: controller}
}

func (r *OrderRouter) RegisterOrderRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /orders", r.controller.GetAllOrder)
	mux.HandleFunc("GET /orders/{id}", r.controller.GetOrderById)
	mux.HandleFunc("POST /orders", r.controller.CreateOrders)
	mux.HandleFunc("PATCH /orders", r.controller.CancelOrder)
	mux.HandleFunc("PATCH /orders", r.controller.UpdateStatusOrder)
	mux.HandleFunc("GET /orders/status", r.controller.GetOrderStatus)

}
