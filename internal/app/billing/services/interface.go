package services

import "videocall/internal/app/billing/schema"

type BillingServiceInterface interface {
	InitialPayment(paymentReq *schema.InitialPaymantRequest) error
	CancelPayment(invoiceId int) error
	VerifyPayment(invoiceId int, status *schema.VerifyPaymentRequest) error
}
