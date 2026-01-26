package services

import "videocall/internal/app/billing/schema"

type BillingServiceInterface interface {
	InitialPayment(paymentReq *schema.InitialPaymantRequest) error
	CancelPayment(userId int) error
	VerifyPayment(userId int, paymentId string) error
}
