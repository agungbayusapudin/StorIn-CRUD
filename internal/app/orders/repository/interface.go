package repository

import "videocall/internal/app/orders/schema"

type OrderRepositoryInterface interface {
	CreateOrder(id int) (*schema.Order, error)
	GetAllOrder() ([]*schema.Order, error)
	UpdateOrder(id int) (*schema.Order, error)
	DeleteOrer(id int) error
	UpdateStatusOrder(id int, status *schema.UpdateStatusOrderRequest) (*schema.Order, error)
	GetOrderById(id int) (*schema.Order, error)
	// user member
	// CreateOrders(id int) (*schema.Order, error)
	// GetAllOrder() ([]*schema.Order, error)
	// GetOrderById(id int) (*schema.Order, error)
	// CancelOrder(id int) error
	//
	// // admin
	// GetOrderStatus(id int) (*schema.Order, error)
	// UpdateOrderStatus(id int, status *schema.UpdateStatusOrderRequest) (*schema.Order, error)
}
