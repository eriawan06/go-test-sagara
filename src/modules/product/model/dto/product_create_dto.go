package dto

import (
	"gopkg.in/guregu/null.v4"
)

type ProductCreateDto struct {
	Name        string      `json:"name" validate:"required,min=1,max=255"`
	Description string      `json:"description" validate:"required,min=1,max=255"`
	Price       int         `json:"price" validate:"required,min=1000,max=99999999"`
	Image       null.String `json:"image" validate:"max=255"`
}
