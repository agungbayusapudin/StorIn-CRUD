package repository

import (
	"database/sql"
	"videocall/internal/app/orders/schema"
)

type orderDetailRepository struct {
	db *sql.DB
}

func NewOrderDetailRepository(db *sql.DB) OrderDetailRepositoryInterface {
	return &orderDetailRepository{db: db}
}

func (repo *orderDetailRepository) CreateOrderDetail(id int) (*schema.OrderDetail, error) {
	return nil, nil
}

func (repo *orderDetailRepository) GetAllOrderDetail() ([]*schema.OrderDetail, error) {
	return nil, nil
}

func (repo *orderDetailRepository) UpdateOrderDetail(id int) (*schema.OrderDetail, error) {
	return nil, nil
}

func (repo *orderDetailRepository) DeleteOrerDetail(id int) error {
	return nil
}

func (repo *orderDetailRepository) UpdateStatusOrderDetail(id int, status *schema.UpdateStatusOrderDetailRequest) (*schema.OrderDetail, error) {
	return nil, nil
}

func (repo *orderDetailRepository) GetOrderByIdDetail(id int) (*schema.OrderDetail, error) {
	return nil, nil
}
