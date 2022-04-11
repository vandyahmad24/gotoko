package category

import (
	"time"
	"vandyahmad/gotoko/helper"
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
