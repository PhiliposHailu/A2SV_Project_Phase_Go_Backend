package models

type BookStatus string

const (
	Available BookStatus = "Available"
	Borrowed  BookStatus = "Borrowed"
)