package repository

import "videocall/internal/app/billing/schema"

type BillingRepositoryInterface interface {
	CreateInvoice(userId int, paymentReq *schema.CreateInvoiceRequest) (*schema.Invoice, error)
	GetInvoice(invoiceId int) (*schema.Invoice, error)
	ListInvoices(userId int) ([]*schema.Invoice, error)
	UpdateInvoiceStatus(invoiceId int, status string) error
	DeleteInvoice(invoiceId int) error
}
