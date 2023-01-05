package payment

import (
	"vandyahmad/newgotoko/helper"
)

type ResponsePayment struct {
	Id   int     `json:"paymentId"`
	Name string  `json:"name"`
	Type string  `json:"type"`
	Logo *string `json:"logo"`
}

type PaymenttListResponse struct {
	Payments []ResponsePayment `json:"payments"`
	Meta     helper.Meta       `json:"meta"`
}
