package routers

import (
	"net/http"
	"videocall/internal/app/billing/controller"
)

type BillingRouter struct {
	ctrl *controller.BillingController
}

func NewBillingRouter(ctrl *controller.BillingController) *BillingRouter {
	return &BillingRouter{ctrl: ctrl}
}

func (router *BillingRouter) RegisterBillingRouter(mux *http.ServeMux) {
	mux.HandleFunc("POST /billing/initial", router.ctrl.InitialPayment)
	mux.HandleFunc("POST /billing/cancel", router.ctrl.CancelPayment)
	mux.HandleFunc("POST /billing/verify", router.ctrl.VerifyPayment)
	mux.HandleFunc("POST /billing/update-status", router.ctrl.UpdatePaymentStatus)
}
