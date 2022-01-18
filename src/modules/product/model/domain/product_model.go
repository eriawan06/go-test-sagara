package domain

import (
	"gopkg.in/guregu/null.v4"
)

type Product struct {
	Id          int         `db:"id"`
	Name        string      `db:"name"`
	Description string      `db:"description"`
	Price       int         `db:"price"`
	Image       null.String `db:"image"`
	CreatedAt   string      `db:"created_at"`
	UpdatedAt   string      `db:"updated_at"`
}
