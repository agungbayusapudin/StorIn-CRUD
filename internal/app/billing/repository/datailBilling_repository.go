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

func (repo *billingDetailRepository) CreateInvoice(paymentReq *schema.CreateInvoiceDetailRequest) error {
	stmt := "INSERT INTO invoice_items (invoice_id, product_name_snapshot, quantity, unit_price, total_price) VALUES ($1, $2, $3, $4)"

	_, err := repo.db.Exec(stmt,
		paymentReq.InvoiceId,
		paymentReq.ProductNameSanapshot,
		paymentReq.Quantity,
		paymentReq.UnitPrice,
		paymentReq.TotalPrice,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *billingDetailRepository) GetInvoice(idDetailInvoice int) (*schema.InvoiceDetail, error) {
	var invoiceDetails schema.InvoiceDetail

	stmt := "SELECT * FROM invoice_items WHERE id = $1"

	rows, err := repo.db.Query(stmt, idDetailInvoice)
	if err != nil {
		return nil, err
	}

	err = rows.Scan(
		&invoiceDetails.Id,
		&invoiceDetails.InvoiceId,
		&invoiceDetails.ProductNameSanapshot,
		&invoiceDetails.Quantity,
		&invoiceDetails.UnitPrice,
		&invoiceDetails.TotalPrice,
	)
	if err != nil {
		return nil, err
	}

	return &invoiceDetails, nil
}

func (repo *billingDetailRepository) DeleteInvoice(invoiceId int) error {
	stmt := "DELETE invoice_items WHERE id = $1"

	_, err := repo.db.Exec(stmt, invoiceId)
	if err != nil {
		return err
	}

	return nil
}
