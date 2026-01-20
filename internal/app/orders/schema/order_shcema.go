package schema

type Order struct {
	ID int `json:"id"`
}

type UpdateStatusOrderRequest struct {
	ID int `json:"id"`
}
