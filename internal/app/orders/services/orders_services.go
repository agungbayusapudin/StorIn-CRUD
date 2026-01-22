package services

import (
	"videocall/internal/app/orders/repository"
	"videocall/internal/app/orders/schema"
)

type orderService struct {
	orderRepo       repository.OrderRepositoryInterface
	orderRepoDetail repository.OrderDetailRepositoryInterface
}

func NewOrderService(repo repository.OrderRepositoryInterface, repoDetail repository.OrderDetailRepositoryInterface) OrderServiceInterface {
	return &orderService{
		orderRepo:       repo,
		orderRepoDetail: repoDetail,
	}
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
