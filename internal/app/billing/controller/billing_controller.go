package controller

import (
	"net/http"
	"strconv"
	"videocall/internal/app/billing/schema"
	"videocall/internal/app/billing/services"
	"videocall/pkg/request"
	"videocall/pkg/response"
)

type BillingController struct {
	service services.BillingServiceInterface
}

func NewBillingController(service services.BillingServiceInterface) *BillingController {
	return &BillingController{service: service}
}

func (ctrl *BillingController) InitialPayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
		return
	}

	var billingReq schema.InitialPaymantRequest

	err := request.DecodeRequest(r, &billingReq)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	err = ctrl.service.InitialPayment(&billingReq)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Payment Initialized"})
}

func (ctrl *BillingController) CancelPayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	err = ctrl.service.CancelPayment(id)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Payment Initialized"})
}

func (ctrl *BillingController) VerifyPayment(w http.ResponseWriter, r *http.Request) {
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

	var invoiceId *schema.VerifyPaymentRequest

	err = request.DecodeRequest(r, &invoiceId)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	err = ctrl.service.VerifyPayment(id, invoiceId)

	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), map[string]string{"status": "Payment Initialized"})
}
