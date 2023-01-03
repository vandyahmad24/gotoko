package cashier

type InputCashier struct {
	Name     string `json:"name" validate:"required"`
	Passcode string `json:"passcode"`
}

type InputPasscode struct {
	Passcode string `json:"passcode" validate:"required"`
}
