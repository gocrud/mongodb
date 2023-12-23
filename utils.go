package mongodb

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

func handleError(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}
	return err
}
