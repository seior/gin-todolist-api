package app

import (
	"database/sql"
	"os"
	"time"
	"todolist-api/helper"
)

func NewDB() *sql.DB {
	dataSource := "root@tcp(localhost:3306)/todolist"

	if os.Getenv("env") == "test" {
		dataSource += "_test"
	} else if os.Getenv("env") != "prod" {
		panic("please set env first")
	}

	db, err := sql.Open("mysql", dataSource)
	helper.PanicIfError(err)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
