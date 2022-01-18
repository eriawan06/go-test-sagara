package domain

type User struct {
	Id        int    `db:"id"`
	Fullname  string `db:"fullname"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
