package request

type DeleteTodolistRequest struct {
	Id int `uri:"id" binding:"required"`
}
