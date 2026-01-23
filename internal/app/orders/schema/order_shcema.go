package schema

type Order struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
