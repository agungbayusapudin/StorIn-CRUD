package repository

import "videocall/internal/app/orders/schema"

type OrderRepositoryInterface interface {
	CreateOrder(id int, totalPrice int) error
	GetAllOrder() ([]*schema.Order, error)
	DeleteOrer(id int) error
	UpdateStatusOrder(id int, status *schema.UpdateStatusOrderRequest) (*schema.Order, error)
	GetOrderById(id int) (*schema.Order, error)
}

type OrderDetailRepositoryInterface interface {
	CreateOrderDetail(id int) (*schema.OrderDetail, error)
	GetAllOrderDetail() ([]*schema.OrderDetail, error)
	UpdateOrderDetail(id int) (*schema.OrderDetail, error)
	DeleteOrerDetail(id int) error
	UpdateStatusOrderDetail(id int, status *schema.UpdateStatusOrderDetailRequest) (*schema.OrderDetail, error)
	GetOrderByIdDetail(id int) (*schema.OrderDetail, error)
}
