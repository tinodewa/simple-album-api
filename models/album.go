package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `json:"title"`
	Artist string             `json:"artist"`
	Price  float32            `json:"price"`
}
