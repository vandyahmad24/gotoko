package category

type InputCategory struct {
	Name string `json:"name" binding:"required"`
}
