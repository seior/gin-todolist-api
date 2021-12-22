package repository

import (
	"context"
	"database/sql"
	"todolist-api/entity"
)

type TodolistRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todolist entity.Todolist) entity.Todolist
	Update(ctx context.Context, tx *sql.Tx, todolist entity.Todolist) entity.Todolist
	FindById(ctx context.Context, tx *sql.Tx, id int) (entity.Todolist, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Todolist
	DeleteById(ctx context.Context, tx *sql.Tx, id int)
	DeleteAll(ctx context.Context, tx *sql.Tx)
}
