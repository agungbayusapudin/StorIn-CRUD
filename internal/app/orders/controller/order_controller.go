package controller

import (
	"net/http"
	"videocall/internal/app/orders/services"
	"videocall/pkg/response"
)

type OrderController struct {
	service services.OrderServiceInterface
}

func NewOrderControllerInterface(service services.OrderServiceInterface) *OrderController {
	return &OrderController{service: service}
}

func (ctrl *OrderController) CreateOrders(w http.ResponseWriter, r *http.Request) {

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Succesfully Created New User"})
}

func (ctrl *OrderController) GetAllOrder(w http.ResponseWriter, r *http.Request) {

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Succesfully Created New User"})
}

func (ctrl *OrderController) GetOrderById(w http.ResponseWriter, r *http.Request) {

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Succesfully Created New User"})
}

func (ctrl *OrderController) CancelOrder(w http.ResponseWriter, r *http.Request) {

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Succesfully Created New User"})
}

func (ctrl *OrderController) GetOrderStatus(w http.ResponseWriter, r *http.Request) {

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Succesfully Created New User"})
}

func (ctrl *OrderController) UpdateStatusOrder(w http.ResponseWriter, r *http.Request) {

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Succesfully Created New User"})
}
