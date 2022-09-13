package models

type TodoModel struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}
