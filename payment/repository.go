package payment

import "gorm.io/gorm"

type Repository interface {
	GetAll(limit, skip int) ([]ResponsePayment, error)
	CountAll() (hasil int64)
	GetById(id int) (ResponsePayment, error)
	Create(payment *Payment) (*ResponsePayment, error)
	Update(id int, payment *Payment) (*ResponsePayment, error)
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

func (r *repository) GetAll(limit, skip int) ([]ResponsePayment, error) {
	var allPayment []ResponsePayment
	rows, err := r.db.Table("payments").Select("id, name, type, logo").Debug().Limit(limit).Offset(skip).Rows()
	if err != nil {
		return allPayment, err
	}

	for rows.Next() {
		var payment Payment
		if err := rows.Scan(&payment.Id, &payment.Name, &payment.Type, &payment.Logo); err != nil {
			return allPayment, err
		}
		var resPayment ResponsePayment
		if payment.Logo == "" {
			resPayment.Logo = nil
		}
		resPayment.Name = payment.Name
		resPayment.Id = payment.Id
		resPayment.Type = payment.Type

		allPayment = append(allPayment, resPayment)

	}

	return allPayment, nil
}

func (r *repository) CountAll() (hasil int64) {
	var payment Payment
	var total int64
	r.db.Model(&payment).Select("count(distinct(id))").Count(&total)

	return total
}

func (r *repository) GetById(id int) (ResponsePayment, error) {

	var allPayment ResponsePayment
	rows, err := r.db.Table("payments").Select("id, name, type, logo").Debug().Where("id = ?", id).Rows()
	if err != nil {
		return allPayment, err
	}

	for rows.Next() {
		var payment Payment
		if err := rows.Scan(&payment.Id, &payment.Name, &payment.Type, &payment.Logo); err != nil {
			return allPayment, err
		}
		if payment.Logo == "" {
			allPayment.Logo = nil
		}
		allPayment.Name = payment.Name
		allPayment.Id = payment.Id
		allPayment.Type = payment.Type

	}

	return allPayment, nil

}

func (r *repository) Create(payment *Payment) (*ResponsePayment, error) {

	err := r.db.Create(payment).Error
	if err != nil {
		return nil, err
	}

	res := ResponsePayment{
		Id:   payment.Id,
		Name: payment.Name,
		Type: payment.Type,
		Logo: nil,
	}

	if payment.Logo != "" {
		res.Logo = &payment.Logo
	}

	return &res, nil
}

func (r *repository) Update(id int, payment *Payment) (*ResponsePayment, error) {
	var oldPayment Payment
	err := r.db.Where("id = ?", id).First(&oldPayment).Error
	if err != nil {
		return nil, err
	}
	//
	oldPayment.Name = payment.Name
	oldPayment.Logo = payment.Logo
	oldPayment.Type = payment.Type
	err = r.db.Save(&oldPayment).Error
	if err != nil {
		return nil, err
	}

	resPayment := ResponsePayment{
		Id:   oldPayment.Id,
		Name: oldPayment.Name,
		Type: oldPayment.Type,
		Logo: &oldPayment.Type,
	}

	return &resPayment, nil
}

func (r *repository) Delete(id int) error {
	
	var payment Payment
	err := r.db.Table("payments").Where("id = ?", id).Delete(&payment).Error
	if err != nil {
		return err
	}
	return nil
}
