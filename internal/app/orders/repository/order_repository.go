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

func (repo *orderRepository) CreateOrder(id int, totalPrice int) error {
	stmt := "INSERT INTO orders (user_id, total_price) VALUES ($1, $2)"

	_, err := repo.db.Exec(stmt, id, totalPrice)
	if err != nil {
		return err
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

		// memasukan ke dalam variabel orderData yang memiliki shcmea order
		err = rows.Scan(&ordersData.ID, &ordersData.UserId, &ordersData.Status, &ordersData.CreatedAt, &ordersData.UpdatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, &ordersData)
	}

	return result, nil
}

func (repo *orderRepository) DeleteOrer(id int) error {
	stmt := "DELETE users WHERE id = $1"

	_, err := repo.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *orderRepository) UpdateStatusOrder(id int, status *schema.UpdateStatusOrderRequest) (*schema.Order, error) {
	var order schema.Order

	stmt := "UPDATE INTO orders SET status = $1. WHERE id = $2 RETURNING *"

	data, err := repo.db.Query(stmt, status.Status, id)
	if err != nil {
		return nil, err
	}

	if data != nil {
		err = data.Scan(&order.ID, &order.UserId, &order.Status, &order.CreatedAt, &order.UpdatedAt)

		if err != nil {
			return nil, err
		}
	}

	return &order, nil
}

func (repo *orderRepository) GetOrderById(id int) (*schema.Order, error) {
	var order schema.Order

	stmt := "SELECT * FROM orders WHERE id = $1"

	rows, err := repo.db.Query(stmt, id)
	if err != nil {
		return nil, err
	}

	if rows != nil {
		err = rows.Scan(&order.ID, &order.UserId, &order.Status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &order, nil
}
