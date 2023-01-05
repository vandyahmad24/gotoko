package payment

import (
	"errors"
)

type Service interface {
	GetAll(limit, skip int) ([]ResponsePayment, error)
	CountAll() (hasil int64)
	GetById(id int) (ResponsePayment, error)
	RegisterPayment(input *RequestPayment) (*ResponsePayment, error)
	UpdatePayment(id int, input *RequestPayment) (*ResponsePayment, error)
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

func (s *service) GetAll(limit, skip int) ([]ResponsePayment, error) {
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

func (s *service) GetById(id int) (ResponsePayment, error) {
	category, err := s.repository.GetById(id)
	if err != nil {
		return category, err
	}
	if category.Name == "" {
		return category, errors.New("Payment Not Found")
	}
	return category, err
}

func (s *service) RegisterPayment(input *RequestPayment) (*ResponsePayment, error) {

	payment := Payment{
		Name: input.Name,
		Type: input.Type,
		Logo: input.Logo,
	}

	newPayment, err := s.repository.Create(&payment)
	if err != nil {
		return nil, err
	}

	return newPayment, nil
}

func (s *service) UpdatePayment(id int, input *RequestPayment) (*ResponsePayment, error) {

	payment := Payment{
		Name: input.Name,
		Type: input.Type,
		Logo: input.Logo,
	}

	newCategory, err := s.repository.Update(id, &payment)
	if err != nil {
		return nil, err
	}
	resPayment := ResponsePayment{
		Id:   newCategory.Id,
		Name: newCategory.Name,
		Type: newCategory.Type,
		Logo: newCategory.Logo,
	}
	return &resPayment, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
