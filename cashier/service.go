package cashier

import "errors"

type Service interface {
	GetAll(limit, skip int) ([]Cashier, error)
	CountAll() (hasil int64)
	GetById(id int) (Cashier, error)
	RegisterCashier(input *InputCashier) (*CashierResponse, error)
	UpdateCashier(id int, input *InputCashier) (*CashierResponse, error)
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

func (s *service) GetAll(limit, skip int) ([]Cashier, error) {
	cashier, err := s.repository.GetAll(limit, skip)
	if err != nil {
		return cashier, err
	}
	return cashier, err
}

func (s *service) CountAll() (hasil int64) {
	total := s.repository.CountAll()
	return total
}

func (s *service) GetById(id int) (Cashier, error) {
	cashier, err := s.repository.GetById(id)
	if err != nil {
		return cashier, err
	}
	if cashier.Name == "" {
		return cashier, errors.New("Cashier Not Found")
	}
	return cashier, err
}

func (s *service) RegisterCashier(input *InputCashier) (*CashierResponse, error) {
	cashier := Cashier{
		Name:     input.Name,
		Passcode: input.Passcode,
	}
	newCashier, err := s.repository.Create(&cashier)
	if err != nil {
		return nil, err
	}
	cashierResponse := CashierResponse{
		Passcode:  newCashier.Passcode,
		Id:        newCashier.Id,
		Name:      newCashier.Name,
		UpdatedAt: newCashier.UpdatedAt,
		CreatedAt: newCashier.CreatedAt,
	}
	return &cashierResponse, nil
}

func (s *service) UpdateCashier(id int, input *InputCashier) (*CashierResponse, error) {
	cashier := Cashier{
		Name:     input.Name,
		Passcode: input.Passcode,
	}
	newCashier, err := s.repository.Update(id, &cashier)
	if err != nil {
		return nil, err
	}
	cashierResponse := CashierResponse{
		Passcode:  newCashier.Passcode,
		Id:        newCashier.Id,
		Name:      newCashier.Name,
		UpdatedAt: newCashier.UpdatedAt,
		CreatedAt: newCashier.CreatedAt,
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
