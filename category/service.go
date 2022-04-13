package category

import "go.uber.org/zap"

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
	zap.L().Info("Creating category service triggered")
	err := s.repo.Create(category)
	if err != nil {
		zap.L().Error("Error creating category (service)", zap.Error(err))
		return err
	}
	return nil
}

func (s *categoryService) List() ([]Category, error) {
	zap.L().Info("Listing category service triggered")
	categories, err := s.repo.List()
	if err != nil {
		zap.L().Error("Error listing category (service)", zap.Error(err))
		return nil, err
	}
	return categories, nil
}
