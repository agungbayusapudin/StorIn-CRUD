package repository

import "videocall/internal/app/orders/schema"

type OrderRepositoryInterface interface {
	CreateOrder(userId int, totalPrice float64) error
	GetAllOrder() ([]*schema.Order, error)
	DeleteOrer(id int) error
	UpdateStatusOrder(id int, status *schema.UpdateStatusOrderRequest) (*schema.Order, error)
	GetOrderById(id int) (*schema.Order, error)
}

type OrderDetailRepositoryInterface interface {
	CreateOrderDetail(orderDetailReq *schema.CreateOrderRequest) error
	GetAllOrderDetail() ([]*schema.OrderDetail, error)
	DeleteOrerDetail(id int) error
	GetOrderByIdDetail(id int) (*schema.OrderDetail, error)
}
