package services

import (
	"videocall/internal/app/billing/repository"
	"videocall/internal/app/billing/schema"
	generateinvoicenumber "videocall/pkg/generateInvoiceNumber"
)

type billingService struct {
	repoBilling       repository.BillingRepositoryInterface
	repoBillingDetail repository.BillingDetailRepositoryInterface
}

func NewBillingService(repoBilling repository.BillingRepositoryInterface, repoBillingDetail repository.BillingDetailRepositoryInterface) BillingServiceInterface {
	return &billingService{
		repoBilling:       repoBilling,
		repoBillingDetail: repoBillingDetail,
	}
}

func (svc *billingService) InitialPayment(paymentReq *schema.InitialPaymantRequest) error {

	// memanggil pkg generate random uuid
	invoiceNumber := generateinvoicenumber.GenerateInvoiceNumber()

	dataInvoiceRquest := &schema.CreateInvoiceRequest{
		CustomerId:    paymentReq.CustomerId,
		Subtotal:      paymentReq.Subtotal,
		TaxtTotal:     paymentReq.TaxtTotal,
		DiscountTotal: paymentReq.DiscountTotal,
		GrandTotal:    paymentReq.GrandTotal,
	}

	// memangil untuk bagian create invoice repository interface
	newId, err := svc.repoBilling.CreateInvoice(invoiceNumber, dataInvoiceRquest)
	if err != nil {
		return err
	}

	dataInvoiceDetailRequest := &schema.CreateInvoiceDetailRequest{
		InvoiceId:            newId,
		ProductNameSanapshot: paymentReq.ProductNameSanapshot,
		Quantity:             paymentReq.Quantity,
		UnitPrice:            paymentReq.UnitPrice,
		TotalPrice:           paymentReq.TotalPrice,
	}

	// memanggil untuk bagain create detail repository inteface
	err = svc.repoBillingDetail.CreateInvoice(dataInvoiceDetailRequest)
	if err != nil {
		return err
	}

	return nil
}

func (svc *billingService) CancelPayment(invoiceId int) error {
	// Delete invoice menggunakan service billing dengan id invoice
	err := svc.repoBilling.DeleteInvoice(invoiceId)
	if err != nil {
		return err
	}

	// delete invoice menggunakan service detail billing dengan id invoice
	err = svc.repoBillingDetail.DeleteInvoice(invoiceId)
	if err != nil {
		return err
	}

	return nil
}

func (svc *billingService) VerifyPayment(invoiceId int, status string) error {
	err := svc.repoBilling.UpdateInvoiceStatus(invoiceId, status)
	if err != nil {
		return err
	}
	return nil
}
