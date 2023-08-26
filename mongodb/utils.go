package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IdsToObjectIds(ids []string) ([]*primitive.ObjectID, error) {
	var objectIds []*primitive.ObjectID
	for _, id := range ids {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIds = append(objectIds, &objectId)
	}
	return objectIds, nil
}
