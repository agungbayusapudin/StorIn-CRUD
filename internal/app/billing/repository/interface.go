package repository

import "videocall/internal/app/billing/schema"

type BillingRepositoryInterface interface {
	CreateInvoice(invoice_number string, paymentReq *schema.CreateInvoiceRequest) (int, error)
	GetInvoice(invoiceId int) (*schema.Invoice, error)
	UpdateInvoiceStatus(invoiceId int, status string) error
	DeleteInvoice(invoiceId int) error
}

type BillingDetailRepositoryInterface interface {
	CreateInvoice(paymentReq *schema.CreateInvoiceDetailRequest) error
	GetInvoice(idDetailInvoice int) (*schema.InvoiceDetail, error)
	DeleteInvoice(invoiceId int) error
}
