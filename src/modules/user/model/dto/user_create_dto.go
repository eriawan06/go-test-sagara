package dto

type UserCreateDto struct {
	Fullname string `json:"fullname" validate:"required,min=1,max=255"`
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}
