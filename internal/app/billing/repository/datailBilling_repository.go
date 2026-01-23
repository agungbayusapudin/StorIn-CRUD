package repository

import (
	"database/sql"
	"videocall/internal/app/billing/schema"
)

type billingDetailRepository struct {
	db *sql.DB
}

func NewBillingDetailRepository(db *sql.DB) BillingDetailRepositoryInterface {
	return &billingDetailRepository{db: db}
}

func (repo *billingDetailRepository) CreateInvoice(userId int, paymentReq *schema.CreateInvoiceRequest) (*schema.Invoice, error) {
	return nil, nil
}

func (repo *billingDetailRepository) GetInvoice(invoiceId int) (*schema.Invoice, error) {
	return nil, nil
}

func (repo *billingDetailRepository) UpdateInvoiceStatus(userId int, status string) error {
	return nil
}

func (repo *billingDetailRepository) DeleteInvoice(invoiceId int) error {
	return nil
}
