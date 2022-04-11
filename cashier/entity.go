package cashier

import (
	"time"
)

type Cashier struct {
	Id        int       `json:"cashierId" gorm:"id"`
	Name      string    `json:"name"  gorm:"name"`
	Passcode  string    `json:"-"  gorm:"passcode"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
