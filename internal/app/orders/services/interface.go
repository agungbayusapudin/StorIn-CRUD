package services

import "videocall/internal/app/orders/schema"

type OrderServiceInterface interface {

	// user member
	CreateOrders(id int) (*schema.Order, error)
	GetAllOrder() ([]*schema.Order, error)
	GetOrderById(id int) (*schema.Order, error)
	CancelOrder(id int) error

	// admin
	GetOrderStatus(id int) (*schema.Order, error)
	UpdateOrderStatus(id int, status *schema.UpdateStatusOrderRequest) (*schema.Order, error)
}

// flow
// 1. User chose product from product list
// 2. User make an order and chos biling status
// 3. User can see the order status
// 4. User can cancle order

// ROLE USER BUT HAVE PRODUCT
// 1. can update status if produt finaly
// 2. can view all orders users buy the product

// ROLE USER MEMBER
// 1. User can order
// 2. user can cancle
// 3. user can add new order
// 4. user can pay
// 5. user can view orders status and order detail
