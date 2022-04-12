package category

type categoryService struct {
	repo CategoryRepository
}

type CategoryService interface {
	Create(category *Category) error
	List() ([]Category, error)
}

// Compile time proof of interface implementation
var _ CategoryService = &categoryService{}

func NewCategoryService(repo CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) Create(category *Category) error {
	err := s.repo.Create(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *categoryService) List() ([]Category, error) {
	categories, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
