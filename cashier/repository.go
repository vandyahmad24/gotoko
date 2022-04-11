package cashier

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(limit, skip int) ([]Cashier, error)
	CountAll() (hasil int64)
	GetById(id int) (Cashier, error)
	Create(cashier *Cashier) (*Cashier, error)
	Update(id int, cashier *Cashier) (*Cashier, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(limit, skip int) ([]Cashier, error) {
	var allCashier []Cashier
	err := r.db.Limit(limit).Offset(skip).Find(&allCashier).Error
	if err != nil {
		return allCashier, err
	}
	return allCashier, nil
}

func (r *repository) CountAll() (hasil int64) {
	var cashier Cashier
	var total int64
	r.db.Model(&cashier).Select("count(distinct(name))").Count(&total)

	return total
}

func (r *repository) GetById(id int) (Cashier, error) {
	var cashier Cashier
	err := r.db.Model(&cashier).Where("id = ?", id).Find(&cashier).Error
	if err != nil {
		return cashier, err
	}
	return cashier, nil
}

func (r *repository) Create(cashier *Cashier) (*Cashier, error) {
	err := r.db.Create(cashier).Error
	if err != nil {
		return cashier, err
	}
	return cashier, nil
}

func (r *repository) Update(id int, cashier *Cashier) (*Cashier, error) {
	var oldCashier Cashier
	err := r.db.Where("id = ?", id).First(&oldCashier).Error
	if err != nil {
		return cashier, err
	}

	oldCashier.Name = cashier.Name
	oldCashier.Passcode = cashier.Passcode
	err = r.db.Save(&oldCashier).Error
	if err != nil {
		return &oldCashier, err
	}

	return &oldCashier, nil
}

func (r *repository) Delete(id int) error {
	var cashier Cashier
	err := r.db.Where("id = ?", id).Delete(&cashier).Error
	if err != nil {
		return err
	}
	return nil
}
