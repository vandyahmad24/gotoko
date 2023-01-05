package payment

type RequestPayment struct {
	Name string `json:"name"  validate:"required"`
	Type string `json:"type"  validate:"required"`
	Logo string `json:"logo"`
}
