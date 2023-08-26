package helpers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenID() string {
	return primitive.NewObjectID().Hex()
}
