package basket

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
	return s.repo.Create(userId)
}

func (s *basketService) UpdateBasket(basket *Basket) error {
	return s.repo.Update(basket)
}

func (s *basketService) GetByUserId(userId uint) (*Basket, error) {
	return s.repo.GetByUserId(userId)
}
