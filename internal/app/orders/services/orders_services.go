package services

import (
	"videocall/internal/app/orders/repository"
	"videocall/internal/app/orders/schema"
)

type orderService struct {
	repo repository.OrderRepositoryInterface
}

func NewOrderService(repo repository.OrderRepositoryInterface) OrderServiceInterface {
	return &orderService{repo: repo}
}

func (serv *orderService) CreateOrders(id int) (*schema.Order, error) {
	return nil, nil
}

func (serv *orderService) GetAllOrder() ([]*schema.Order, error) {
	return nil, nil
}

func (serv *orderService) GetOrderById(id int) (*schema.Order, error) {
	return nil, nil
}

func (serv *orderService) CancelOrder(id int) error {
	return nil
}

func (serv *orderService) GetOrderStatus(id int) (*schema.Order, error) {
	return nil, nil
}

func (serv *orderService) UpdateOrderStatus(id int, status *schema.UpdateStatusOrderRequest) (*schema.Order, error) {
	return nil, nil
}
