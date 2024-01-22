package mongodb

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

func handleError(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	} else if errors.Is(err, mongo.ErrNilDocument) {
		return nil
	}
	return err
}

func HexID() string {
	return primitive.NewObjectID().Hex()
}

func IsObjectID(id string) bool {
	_, err := primitive.ObjectIDFromHex(id)
	return err == nil
}

func ObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}
