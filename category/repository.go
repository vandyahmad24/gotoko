package category

import "gorm.io/gorm"

type Repository interface {
	GetAll(limit, skip int) ([]Category, error)
	CountAll() (hasil int64)
	GetById(id int) (Category, error)
	Create(category *Category) (*Category, error)
	Update(id int, category *Category) (*Category, error)
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

func (r *repository) GetAll(limit, skip int) ([]Category, error) {
	var allCategory []Category
	err := r.db.Limit(limit).Offset(skip).Find(&allCategory).Error
	if err != nil {
		return allCategory, err
	}
	return allCategory, nil
}

func (r *repository) CountAll() (hasil int64) {
	var category Category
	var total int64
	r.db.Model(&category).Select("count(distinct(name))").Count(&total)

	return total
}

func (r *repository) GetById(id int) (Category, error) {
	var category Category
	err := r.db.Model(&category).Where("id = ?", id).Find(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repository) Create(category *Category) (*Category, error) {
	err := r.db.Create(category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repository) Update(id int, category *Category) (*Category, error) {
	var oldCategory Category
	err := r.db.Where("id = ?", id).First(&oldCategory).Error
	if err != nil {
		return category, err
	}

	oldCategory.Name = category.Name
	err = r.db.Save(&oldCategory).Error
	if err != nil {
		return &oldCategory, err
	}

	return &oldCategory, nil
}

func (r *repository) Delete(id int) error {
	var category Category
	err := r.db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}
