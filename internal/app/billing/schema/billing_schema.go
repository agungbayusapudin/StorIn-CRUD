package schema

type Invoice struct {
	Id int `json:"id"`
}

type InvoiceDetail struct {
	Id int `json:"id"`
}

type CreateInvoiceRequest struct {
	UserId int     `json:"user_id"`
	Method string  `json:"method"`
	Amount float64 `json:"amount"`
}
