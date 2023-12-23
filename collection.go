package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	name string
	col  *mongo.Collection
	db   *Database
}

func (c *Collection) Query() *Query {
	return &Query{
		c:   c,
		ctx: context.Background(),
	}
}

func (c *Collection) InsertOne(data any) error {
	_, err := c.col.InsertOne(context.Background(), data)
	return handleError(err)
}

func (c *Collection) InsertMany(data []any) error {
	_, err := c.col.InsertMany(context.Background(), data)
	return handleError(err)
}
