package dto

type UserResponseDto struct {
	Id        int    `json:"id"`
	Fullname  string `json:"fullname"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserResponseWithPasswordDto struct {
	Id        int    `json:"id"`
	Fullname  string `json:"fullname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
