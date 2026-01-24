package schema

type Invoice struct {
	Id            int     `json:"id"`
	InvoiceNumber string  `json:"invoice_number"`
	CustomerId    int     `json:"customer_id"`
	IssueDate     string  `json:"issue_date"`
	DueDate       string  `json:"due_date"`
	Subtotal      float64 `json:"subtotal"`
	TaxtTotal     float64 `json:"taxt_total"`
	DiscountTotal float64 `json:"discount_total"`
	GrandTotal    float64 `json:"grand_total"`
	Satus         string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
}

type InvoiceDetail struct {
	Id                   int     `json:"id"`
	InvoiceId            int     `json:"product_id"`
	ProductNameSanapshot string  `json:"product_name_snapshot"`
	Quantity             int     `json:"quantity"`
	UnitPrice            float64 `json:"unit_price"`
	TotalPrice           float64 `json:"total_price"`
}

type CreateInvoiceRequest struct {
	CustomerId    int     `json:"customer_id"`
	Subtotal      float64 `json:"subtotal"`
	TaxtTotal     float64 `json:"taxt_total"`
	DiscountTotal float64 `json:"discount_total"`
	GrandTotal    float64 `json:"grand_total"`
}
