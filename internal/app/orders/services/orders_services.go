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

func (serv *orderService) CreateOrders(orderRequest *schema.CreateOrderRequest) error {
	// melakukan perkalian(*) qty dengan harga product = totalPrice
	totalPrice := float64(orderRequest.Quantity) * orderRequest.PriceAtPurchase

	err := serv.orderRepo.CreateOrder(orderRequest.UserId, float64(totalPrice))
	if err != nil {
		return err
	}

	err = serv.orderRepoDetail.CreateOrderDetail(orderRequest)
	if err != nil {
		return err
	}

	return nil
}

func (serv *orderService) GetAllOrder() ([]*schema.Order, error) {
	data, err := serv.orderRepo.GetAllOrder()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (serv *orderService) GetOrderById(id int) (*schema.Order, error) {
	data, err := serv.orderRepo.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (serv *orderService) CancelOrder(id int, status *schema.UpdateStatusOrderRequest) (*schema.Order, error) {
	// melakukan update status ordernya
	data, err := serv.orderRepo.UpdateStatusOrder(id, status)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// untuk admin
// tambahkan agar lengkap dalam mengembalikan datanya
func (serv *orderService) UpdateOrderStatus(id int, status *schema.UpdateStatusOrderRequest) (*schema.Order, error) {
	data, err := serv.orderRepo.UpdateStatusOrder(id, status)
	if err != nil {
		return nil, err
	}

	return data, nil
}
