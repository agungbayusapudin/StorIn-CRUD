package services

import (
	"videocall/internal/app/billing/repository"
	"videocall/internal/app/billing/schema"
)

type billingService struct {
	repo repository.BillingRepositoryInterface
}

func NewBillingService(repo repository.BillingRepositoryInterface) BillingServiceInterface {
	return &billingService{repo: repo}
}

func (svc *billingService) InitialPayment(userId int, paymentReq *schema.CreateInvoiceRequest) error {
	return nil
}

func (svc *billingService) CancelPayment(userId int) error {
	return nil
}

func (svc *billingService) VerifyPayment(userId int, paymentId string) error {
	return nil
}

func (svc *billingService) UpdatePaymentStatus(userId int, status string) error {
	return nil
}
