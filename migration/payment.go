package migration

import "time"

type Payment struct {
	Id        int     `gorm:"column:id"`
	Name      string  `gorm:"size:255"`
	Type      string  `gorm:"size:255"`
	Logo      *string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
