package schema

type Order struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type OrderDetail struct {
	ID              int     `json:"id"`
	OrderId         int     `json:"order_id"`
	ProductId       int     `json:"product_id"`
	Quantity        int     `json:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase"`
}

type CreateOrderRequest struct {
	UserId          int     `json:"user_id"`
	Status          bool    `json:"status"`
	OrderId         int     `json:"order_id"`
	ProductId       int     `json:"product_id"`
	Quantity        int     `json:"quantity"`
	PriceAtPurchase float64 `json:"price_at_purchase"`
}

type UpdateStatusOrderRequest struct {
	Status bool `json:"status"`
}

type UpdateStatusOrderDetailRequest struct {
	ID int `json:"id"`
}
