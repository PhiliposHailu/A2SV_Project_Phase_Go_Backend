package models

type Task struct {
	Title string `json:"title"`
	Description string `json:"description"`
	DueDate string `json:"due_date"`
	Status string `json:"status"`
}