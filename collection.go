package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	name string
	col  *mongo.Collection
	db   *Database
}

func (c *Collection) Drop() error {
	return c.col.Drop(context.Background())
}

func (c *Collection) Query() *Query {
	return &Query{
		c:   c,
		ctx: context.Background(),
	}
}

func (c *Collection) Aggregate() *Aggregation {
	return &Aggregation{
		c:   c,
		ctx: context.Background(),
	}
}

func (c *Collection) InsertOne(data any) (primitive.ObjectID, error) {
	result, err := c.col.InsertOne(context.Background(), data)
	if err != nil {
		return primitive.NilObjectID, handleError(err)
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (c *Collection) InsertMany(data []any) ([]primitive.ObjectID, error) {
	result, err := c.col.InsertMany(context.Background(), data)
	if err != nil {
		return nil, handleError(err)
	}
	var ids []primitive.ObjectID
	for _, id := range result.InsertedIDs {
		ids = append(ids, id.(primitive.ObjectID))
	}
	return ids, nil
}
