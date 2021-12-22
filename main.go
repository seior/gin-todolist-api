package main

import (
	_ "github.com/go-sql-driver/mysql"
	"todolist-api/helper"
)

func main() {
	server := InitializeServer()

	err := server.Run()
	helper.PanicIfError(err)
}
