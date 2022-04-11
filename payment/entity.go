package payment

import "time"

type Payment struct {
	Id        int       `json:"paymentId" gorm:"id"`
	Name      string    `json:"name"  gorm:"name"`
	Type      string    `json:"type"  gorm:"type"`
	Logo      string    `json:"logo"  gorm:"logo"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
