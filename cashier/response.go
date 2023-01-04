package cashier

import (
	"time"
	"vandyahmad/newgotoko/helper"
)

type CashierListResponse struct {
	Cashiers []Cashier   `json:"cashiers"`
	Meta     helper.Meta `json:"meta"`
}

type CashierResponse struct {
	Passcode  string    `json:"passcode"`
	Id        int       `json:"cashierId"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type CashierPasscode struct {
	Passcode string `json:"passcode"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
