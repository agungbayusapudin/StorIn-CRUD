package services

import (
	"videocall/internal/app/billing/repository"
	"videocall/internal/app/billing/schema"
)

type billingDetailService struct {
	repo repository.BillingDetailRepositoryInterface
}

func NewBillingDetailRepository(repo repository.BillingDetailRepositoryInterface) BillingDetailServiceInterface {
	return &billingDetailService{repo: repo}
}

func (svc *billingDetailService) InitialPayment(userId int, paymentReq *schema.CreateInvoiceRequest) error {
	return nil
}

func (svc *billingDetailService) CancelPayment(user_id int) error {
	return nil
}

func (svc *billingDetailService) VerifyPayment(user_id int, paymentId string) error {
	return nil
}

func (svc *billingDetailService) UpdatePaymentStatus(invoiceId int, status string) error {
	return nil
}
