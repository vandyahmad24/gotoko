package category

import "errors"

type Service interface {
	GetAll(limit, skip int) ([]Category, error)
	CountAll() (hasil int64)
	GetById(id int) (Category, error)
	RegisterCategory(input *InputCategory) (*CategoryResponse, error)
	UpdateCategory(id int, input *InputCategory) (*CategoryResponse, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll(limit, skip int) ([]Category, error) {
	category, err := s.repository.GetAll(limit, skip)
	if err != nil {
		return category, err
	}
	return category, err
}

func (s *service) CountAll() (hasil int64) {
	total := s.repository.CountAll()
	return total
}

func (s *service) GetById(id int) (Category, error) {
	category, err := s.repository.GetById(id)
	if err != nil {
		return category, err
	}
	if category.Name == "" {
		return category, errors.New("Category Not Found")
	}
	return category, err
}

func (s *service) RegisterCategory(input *InputCategory) (*CategoryResponse, error) {
	category := Category{
		Name: input.Name,
	}
	newCategory, err := s.repository.Create(&category)
	if err != nil {
		return nil, err
	}
	cashierResponse := CategoryResponse{
		Id:        newCategory.Id,
		Name:      newCategory.Name,
		UpdatedAt: newCategory.UpdatedAt,
		CreatedAt: newCategory.CreatedAt,
	}
	return &cashierResponse, nil
}

func (s *service) UpdateCategory(id int, input *InputCategory) (*CategoryResponse, error) {
	category := Category{
		Name: input.Name,
	}
	newCategory, err := s.repository.Update(id, &category)
	if err != nil {
		return nil, err
	}
	cashierResponse := CategoryResponse{
		Id:        newCategory.Id,
		Name:      newCategory.Name,
		UpdatedAt: newCategory.UpdatedAt,
		CreatedAt: newCategory.CreatedAt,
	}
	return &cashierResponse, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
