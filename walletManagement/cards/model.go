package cards

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name"`
}
