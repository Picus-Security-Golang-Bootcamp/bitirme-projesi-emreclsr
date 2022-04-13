package basket

import "go.uber.org/zap"

type basketService struct {
	repo BasketRepository
}

type BasketService interface {
	CreateBasket(userId uint) error
	UpdateBasket(basket *Basket) error
	GetByUserId(userId uint) (*Basket, error)
}

// Compile time proof of interface implementation
var _ BasketService = &basketService{}

func NewBasketService(repo BasketRepository) BasketService {

	return &basketService{repo: repo}
}

func (s *basketService) CreateBasket(userId uint) error {
	zap.L().Info("Create basket service triggered")
	return s.repo.Create(userId)
}

func (s *basketService) UpdateBasket(basket *Basket) error {
	zap.L().Info("Update basket service triggered")
	return s.repo.Update(basket)
}

func (s *basketService) GetByUserId(userId uint) (*Basket, error) {
	zap.L().Info("Get basket service triggered")
	return s.repo.GetByUserId(userId)
}
