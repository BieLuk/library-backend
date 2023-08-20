package dto

import "github.com/google/uuid"

type GetBooksResponse struct {
	Books []GetBookResponse `json:"books"`
}

type GetBookResponse struct {
	ID          *uuid.UUID `json:"id"`
	Name        string     `json:"name"`
	Author      string     `json:"author"`
	ISBN        string     `json:"isbn"`
	Description *string    `json:"description,omitempty"`
}

type CreateBookRequest struct {
	Name        string  `json:"name" validate:"required,max=255"`
	Author      string  `json:"author" validate:"required,max=255"`
	ISBN        string  `json:"isbn" validate:"required,len=13,number"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=255"`
}

type CreateBookResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateBookRequest struct {
	Name        string  `json:"name" validate:"required,max=255"`
	Author      string  `json:"author" validate:"required,max=255"`
	ISBN        string  `json:"isbn" validate:"required,len=13,number"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=255"`
}
