package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Diario struct {
	ID        primitive.ObjectID `bson:"id,omitempty"`
	Id_user   int                `bson:"id_user"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created"`
}
