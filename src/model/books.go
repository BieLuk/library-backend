package model

type Book struct {
	DBEntity    `bson:",inline"`
	Name        string  `bson:"name"`
	Author      string  `bson:"author"`
	ISBN        string  `bson:"isbn"`
	Description *string `bson:"description,omitempty"`
}
