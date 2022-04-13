package order

import "go.uber.org/zap"

type orderService struct {
	repo OrderRepository
}

type OrderService interface {
	Create(order *Order) error
	Get(id uint) (*Order, error)
	List(userID uint) ([]Order, error)
	Update(order *Order) error
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderService{repo: repo}
}

// Compile time proof of interface implementation
var _ OrderService = &orderService{}

func (s *orderService) Create(order *Order) error {
	zap.L().Info("Create order service triggered")
	err := s.repo.Create(order)
	if err != nil {
		return err
	}
	return nil
}

func (s *orderService) Get(id uint) (*Order, error) {
	zap.L().Info("Get order service triggered")
	order, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *orderService) List(userID uint) ([]Order, error) {
	zap.L().Info("List order service triggered")
	orders, err := s.repo.List(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *orderService) Update(order *Order) error {
	zap.L().Info("Update order service triggered")
	err := s.repo.Update(order)
	if err != nil {
		return err
	}
	return nil
}
