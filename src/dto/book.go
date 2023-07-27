package dto

import "github.com/google/uuid"

type GetBooksResponse struct {
	Books []GetBookResponse `json:"books"`
}

type GetBookResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	Description *string   `json:"description,omitempty"`
}

type CreateBookRequest struct {
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	ISBN        string  `json:"isbn"`
	Description *string `json:"description,omitempty"`
}

type CreateBookResponse struct {
	ID uuid.UUID `json:"id"`
}
