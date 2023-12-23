package mongodb

import (
	"context"
)

type Query struct {
	c      *Collection
	ctx    context.Context
	filter interface{}
}

func (q *Query) Filter(filter interface{}) *Query {
	q.filter = filter
	return q
}

func (q *Query) Context(ctx context.Context) *Query {
	q.ctx = ctx
	return q
}

func (q *Query) FindOne(data any) error {
	err := q.c.col.FindOne(q.ctx, q.filter).Decode(data)
	return handleError(err)
}

func (q *Query) Find(data any) error {
	cursor, err := q.c.col.Find(q.ctx, q.filter)
	if err != nil {
		return handleError(err)
	}
	defer cursor.Close(q.ctx)
	return cursor.All(q.ctx, data)
}

func (q *Query) Count() (int64, error) {
	return q.c.col.CountDocuments(q.ctx, q.filter)
}

func (q *Query) DeleteOne() error {
	_, err := q.c.col.DeleteOne(q.ctx, q.filter)
	return handleError(err)
}

func (q *Query) DeleteMany() error {
	_, err := q.c.col.DeleteMany(q.ctx, q.filter)
	return handleError(err)
}

func (q *Query) UpdateOne(data any) error {
	_, err := q.c.col.UpdateOne(q.ctx, q.filter, data)
	return handleError(err)
}

func (q *Query) UpdateMany(data any) error {
	_, err := q.c.col.UpdateMany(q.ctx, q.filter, data)
	return handleError(err)
}

func (q *Query) ReplaceOne(data any) error {
	_, err := q.c.col.ReplaceOne(q.ctx, q.filter, data)
	return handleError(err)
}
