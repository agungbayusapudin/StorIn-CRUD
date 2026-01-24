package repository

import (
	"database/sql"
	"time"
	"videocall/internal/app/billing/schema"
)

type BillingRepository struct {
	db *sql.DB
}

func NewBillingRepository(db *sql.DB) BillingRepositoryInterface {
	return &BillingRepository{db: db}
}

func (repo *BillingRepository) CreateInvoice(userId int, invoice_number string, paymentReq *schema.CreateInvoiceRequest) error {
	stmt := "INSERT INTO invoices (invoice_number, customer_id, issue_date, due_date, subtotal, taxt_total, discount_total, grand_total, status, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	_, err := repo.db.Exec(stmt,
		invoice_number,
		paymentReq.CustomerId,
		time.Now(),
		time.Now().AddDate(0, 0, 1),
		paymentReq.Subtotal,
		paymentReq.TaxtTotal,
		paymentReq.DiscountTotal,
		paymentReq.GrandTotal,
		"UNPAID",
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}
func (repo *BillingRepository) GetInvoice(invoiceId int) (*schema.Invoice, error) {
	var data schema.Invoice

	stmt := "SELECT * FROM invoices WHERE id = $1"
	rows, err := repo.db.Query(stmt, invoiceId)
	if err != nil {
		return nil, err
	}

	err = rows.Scan(
		&data.Id,
		&data.InvoiceNumber,
		&data.CustomerId,
		&data.IssueDate,
		&data.DueDate,
		&data.Subtotal,
		&data.TaxtTotal,
		&data.DiscountTotal,
		&data.GrandTotal,
		&data.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo *BillingRepository) UpdateInvoiceStatus(invoiceId int, status string) error {
	stmt := "UPDATE invoice SET status = $1 WHERE id = $2"

	_, err := repo.db.Exec(stmt, status, invoiceId)
	if err != nil {
		return err
	}

	return nil
}

func (repo *BillingRepository) DeleteInvoice(invoiceId int) error {
	stmt := "DELETE invoice WHERE id = $1"

	_, err := repo.db.Exec(stmt, invoiceId)
	if err != nil {
		return err
	}

	return nil
}
