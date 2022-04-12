package product

type productService struct {
	repo ProductRepository
}

type ProductService interface {
	Create(product *Product) error
	Search(word string) ([]Product, error)
	Delete(id uint) error
	Update(product *Product) error
	List() ([]Product, error)
	Get(id uint) (*Product, error)
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

// Compile time proof of interface implementation
var _ ProductService = &productService{}

func (s *productService) Create(product *Product) error {
	err := s.repo.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) Search(word string) ([]Product, error) {
	products, err := s.repo.Search(word)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) Update(product *Product) error {
	err := s.repo.Update(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) List() ([]Product, error) {
	products, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) Get(id uint) (*Product, error) {
	product, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
