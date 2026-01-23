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

func (repo *orderDetailRepository) CreateOrderDetail(orderDetailReq *schema.CreateOrderRequest) error {
	stmt := "INSERT INTO order_items (order_id, product_id, quantity, price_at_purchase) VALUES ($1, 2$, 3$, 4$)"

	_, err := repo.db.Exec(stmt, orderDetailReq.OrderId, orderDetailReq.ProductId, orderDetailReq.Quantity, orderDetailReq.PriceAtPurchase)
	if err != nil {
		return nil
	}
	return nil
}

func (repo *orderDetailRepository) GetAllOrderDetail() ([]*schema.OrderDetail, error) {
	var orderDetail []*schema.OrderDetail

	stmt := "SELECT * FROM order_items"

	rows, err := repo.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var orderDetailData schema.OrderDetail

		err = rows.Scan(&orderDetailData.ID, &orderDetailData.OrderId, &orderDetailData.ProductId, &orderDetailData.Quantity, &orderDetailData.PriceAtPurchase)
		if err != nil {
			return nil, err
		}

		orderDetail = append(orderDetail, &orderDetailData)
	}

	return orderDetail, nil
}

func (repo *orderDetailRepository) DeleteOrerDetail(id int) error {
	stmt := "DELETE users WHERE id = $1"

	_, err := repo.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *orderDetailRepository) GetOrderByIdDetail(id int) (*schema.OrderDetail, error) {
	var orderDetail schema.OrderDetail

	stmt := "SELECT * FROM order_items WHERE id = $1"
	rows, err := repo.db.Query(stmt, id)
	if err != nil {
		return nil, err
	}

	err = rows.Scan(&orderDetail.ID, &orderDetail.OrderId, &orderDetail.ProductId, &orderDetail.Quantity, &orderDetail.PriceAtPurchase)
	if err != nil {
		return nil, err
	}

	return &orderDetail, nil
}
