package controller

import "github.com/google/uuid"

type GetBooksResponse struct {
	Books []GetBookResponse
}

type GetBookResponse struct {
	ID          uuid.UUID
	Name        string
	Author      string
	ISBN        string
	Description string
}
