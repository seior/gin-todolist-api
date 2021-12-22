package request

type FindByIdTodolistRequest struct {
	Id int `uri:"id" binding:"required"`
}
