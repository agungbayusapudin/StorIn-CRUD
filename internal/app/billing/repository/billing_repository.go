package repository

import (
	"database/sql"
	"videocall/internal/app/billing/schema"
)

type BillingRepository struct {
	db *sql.DB
}

func NewBillingRepository(db *sql.DB) BillingRepositoryInterface {
	return &BillingRepository{db: db}
}

func (repo *BillingRepository) CreateInvoice(userId int, paymentReq *schema.CreateInvoiceRequest) (*schema.Invoice, error) {
	return nil, nil
}
func (repo *BillingRepository) GetInvoice(invoiceId int) (*schema.Invoice, error) {
	return nil, nil
}

func (repo *BillingRepository) UpdateInvoiceStatus(invoiceId int, status string) error {
	return nil
}

func (repo *BillingRepository) DeleteInvoice(invoiceId int) error {
	return nil
}
