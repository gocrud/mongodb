package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	name string
	db   *mongo.Database
	mg   *MongoDB
}

func (d *Database) Collection(name string) *Collection {
	return &Collection{
		name: name,
		col:  d.db.Collection(name),
		db:   d,
	}
}

func (d *Database) Drop() error {
	return d.db.Drop(context.Background())
}
