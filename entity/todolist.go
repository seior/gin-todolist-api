package entity

type Todolist struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
