package app

import (
	"github.com/eriawan06/go-test-sagara/src/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB() *sqlx.DB {
	DB, err := sqlx.Open("mysql", "root@tcp(localhost:3306)/test_sagara")
	helper.PanicIfError(err)
	err = DB.Ping()
	helper.PanicIfError(err)
	return DB
}
