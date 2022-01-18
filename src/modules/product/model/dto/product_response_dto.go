package dto

import (
	"gopkg.in/guregu/null.v4"
)

type ProductResponseDto struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       int         `json:"price"`
	Image       null.String `json:"image"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}
