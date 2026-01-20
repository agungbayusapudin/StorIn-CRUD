package controller

import (
	"net/http"
	"videocall/internal/app/billing/services"
	"videocall/pkg/response"
)

type BillingController struct {
	service services.BillingServiceInterface
}

func NewBillingController(service services.BillingServiceInterface) *BillingController {
	return &BillingController{service: service}
}

func (ctrl *BillingController) InitialPayment(w http.ResponseWriter, r *http.Request) {
	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Payment Initialized"})
}

func (ctrl *BillingController) CancelPayment(w http.ResponseWriter, r *http.Request) {
	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Payment Initialized"})
}

func (ctrl *BillingController) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Payment Initialized"})
}

func (ctrl *BillingController) UpdatePaymentStatus(w http.ResponseWriter, r *http.Request) {
	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Payment Initialized"})
}
