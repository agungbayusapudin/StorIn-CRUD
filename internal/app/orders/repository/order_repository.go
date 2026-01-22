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

func (repo *orderRepository) CreateOrder(id int) error {
	stmt := "INSERT INTO orders (user_id, total_price) VALUES ($1, $2)"

	_, err := repo.db.Exec(stmt, id, id)
	if err != nil {
		return err
	}

	if err != nil {
		return nil
	}

	return nil
}

// catatan [Agung dev] : Sebaiknya menggunakan pagnation jangan langsung ambil semuanya
func (repo *orderRepository) GetAllOrder() ([]*schema.Order, error) {
	var result []*schema.Order

	stmt := "SELECT * FROM orders"
	rows, err := repo.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ordersData schema.Order

		order, err := rows.Scan(ordersData.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, order)
	}

	return result, nil
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
