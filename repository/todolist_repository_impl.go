package repository

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"todolist-api/entity"
	"todolist-api/helper"
)

type TodolistRepositoryImpl struct {
}

func NewTodolistRepository() TodolistRepository {
	return &TodolistRepositoryImpl{}
}

func (repository *TodolistRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todolist entity.Todolist) entity.Todolist {
	SQL := "insert into todolist (name, author) value (?, ?)"

	result, err := tx.ExecContext(ctx, SQL, todolist.Name, todolist.Author)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	todolist.Id = int(id)

	return todolist
}

func (repository *TodolistRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todolist entity.Todolist) entity.Todolist {
	SQL := "update todolist set name = ?, author = ? where id = ?"

	_, err := tx.ExecContext(ctx, SQL, todolist.Name, todolist.Author, todolist.Id)
	helper.PanicIfError(err)

	return todolist
}

func (repository *TodolistRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Todolist, error) {
	SQL := "select id, name, author from todolist where id = ?"

	result, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer result.Close()

	todolist := entity.Todolist{}
	if result.Next() {
		err := result.Scan(&todolist.Id, &todolist.Name, &todolist.Author)
		helper.PanicIfError(err)

		return todolist, nil
	}

	return todolist, errors.New("todolist not found")
}

func (repository *TodolistRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Todolist {
	SQL := "select id, name, author from todolist"

	result, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer result.Close()

	var todolists []entity.Todolist
	for result.Next() {
		todolist := entity.Todolist{}

		err := result.Scan(&todolist.Id, &todolist.Name, &todolist.Author)
		helper.PanicIfError(err)

		todolists = append(todolists, todolist)
	}

	return todolists
}

func (repository *TodolistRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "delete from todolist where id = ?"

	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (repository *TodolistRepositoryImpl) DeleteAll(ctx context.Context, tx *sql.Tx) {
	if os.Getenv("env") == "test" {
		SQL := "delete from todolist"

		_, err := tx.ExecContext(ctx, SQL)
		helper.PanicIfError(err)
	} else {
		panic("you must be set env to test for deleteAll function")
	}
}
