package models

type ModelTodo struct {
	UserID    int64  `json:"userId"`
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
