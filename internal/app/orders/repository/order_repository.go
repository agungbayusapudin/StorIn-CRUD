package repository

import (
	"database/sql"
	"videocall/internal/app/orders/schema"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepositoryInterface {
	return &orderRepository{db: db}
}

func (repo *orderRepository) CreateOrder(id int) (*schema.Order, error) {
	stmt := "INSERT INTO orders ()"

	_, err := repo.db.Exec(stmt, id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (repo *orderRepository) GetAllOrder() ([]*schema.Order, error) {

	return nil, nil
}

func (repo *orderRepository) UpdateOrder(id int) (*schema.Order, error) {
	return nil, nil
}

func (repo *orderRepository) DeleteOrer(id int) error {
	return nil
}

func (repo *orderRepository) UpdateStatusOrder(id int, status *schema.UpdateStatusOrderRequest) (*schema.Order, error) {
	return nil, nil
}

func (repo *orderRepository) GetOrderById(id int) (*schema.Order, error) {
	return nil, nil
}
