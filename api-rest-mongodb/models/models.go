package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Create struct models
type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitemty"`
	Isbn   string             `json:"isbn,omitempty" bson:"isbn,omitemty"`
	Title  string             `json:"title" bson:"_title,omitemty"`
	Author *Author            `json:"author" bson:"author,omitemty"`
}

type Author struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
