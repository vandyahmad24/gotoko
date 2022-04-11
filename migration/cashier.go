package migration

import "time"

type Cashier struct {
	Id        int    `gorm:"column:id"`
	Name      string `gorm:"size:255"`
	Passcode  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
