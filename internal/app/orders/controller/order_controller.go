package controller

import (
	"net/http"
	"strconv"
	"videocall/internal/app/orders/schema"
	"videocall/internal/app/orders/services"
	"videocall/pkg/request"
	"videocall/pkg/response"
)

type OrderController struct {
	service services.OrderServiceInterface
}

func NewOrderControllerInterface(service services.OrderServiceInterface) *OrderController {
	return &OrderController{service: service}
}

func (ctrl *OrderController) CreateOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
		return
	}

	var reqOrder schema.CreateOrderRequest

	err := request.DecodeRequest(r, &reqOrder)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	err = ctrl.service.CreateOrders(&reqOrder)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Succesfully Created New User"})
}

func (ctrl *OrderController) GetAllOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
		return
	}

	data, err := ctrl.service.GetAllOrder()
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), data)
}

func (ctrl *OrderController) GetOrderById(w http.ResponseWriter, r *http.Request) {
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

	data, err := ctrl.service.GetOrderById(id)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), data)
}

func (ctrl *OrderController) CancelOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	var status schema.UpdateStatusOrderRequest

	err = request.DecodeRequest(r, &status)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	data, err := ctrl.service.CancelOrder(id, &status)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), data)
}

func (ctrl *OrderController) UpdateStatusOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	var status schema.UpdateStatusOrderRequest

	err = request.DecodeRequest(r, &status)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	data, err := ctrl.service.UpdateOrderStatus(id, &status)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), data)
}
