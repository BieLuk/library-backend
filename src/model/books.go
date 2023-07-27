package model

type Book struct {
	DBEntity
	Name        string
	Author      string
	ISBN        string
	Description *string
}
