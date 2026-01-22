package schema

type Order struct {
	ID int `json:"id"`
}

type OrderDetail struct {
	ID int `json:"id"`
}

type UpdateStatusOrderRequest struct {
	ID int `json:"id"`
}

type UpdateStatusOrderDetailRequest struct {
	ID int `json:"id"`
}
