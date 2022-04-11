package cashier

type InputCashier struct {
	Name     string `json:"name" binding:"required"`
	Passcode string `json:"passcode" binding:"required,numeric,min=6,max=6"`
}

type InputPasscode struct {
	Passcode string `json:"passcode" binding:"required,numeric,min=6,max=6"`
}
