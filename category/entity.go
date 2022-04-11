package category

import "time"

type Category struct {
	Id        int       `json:"categoryId" gorm:"id"`
	Name      string    `json:"name"  gorm:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
