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

func (c *Collection) Name() string {
	return c.name
}

func (c *Collection) Drop() error {
	return c.col.Drop(context.Background())
}

func (c *Collection) Query() *Query {
	return &Query{
		c:      c,
		filter: make(map[string]interface{}),
		ctx:    context.Background(),
	}
}

func (c *Collection) Aggregate() *Aggregation {
	return &Aggregation{
		c:   c,
		ctx: context.Background(),
	}
}

func (c *Collection) InsertOne(ctx context.Context, data any) (any, error) {
	result, err := c.col.InsertOne(ctx, data)
	if err != nil {
		return nil, handleError(err)
	}
	return result.InsertedID, nil
}

func (c *Collection) InsertMany(ctx context.Context, data []any) ([]any, error) {
	result, err := c.col.InsertMany(ctx, data)
	if err != nil {
		return nil, handleError(err)
	}
	return result.InsertedIDs, nil
}
