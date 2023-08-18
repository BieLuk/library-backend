package model

type BookStatus string

var (
	BookStatusActive  BookStatus = "ACTIVE"
	BookStatusDeleted BookStatus = "DELETED"
)

type Book struct {
	DBEntity    `bson:",inline"`
	Name        string     `bson:"name,omitempty"`
	Author      string     `bson:"author,omitempty"`
	ISBN        string     `bson:"isbn,omitempty"`
	Description *string    `bson:"description,omitempty"`
	Status      BookStatus `bson:"status,omitempty"`
}
