package category

import (
	"time"
	"vandyahmad/newgotoko/helper"
)

type CategoryListResponse struct {
	Category []Category  `json:"categories"`
	Meta     helper.Meta `json:"meta"`
}

type CategoryResponse struct {
	Id        int       `json:"cashierId"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
