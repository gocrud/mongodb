package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Query struct {
	c      *Collection
	ctx    context.Context
	filter interface{}
	sort   interface{}
	limit  int64
	skip   int64
	proj   interface{}
	//collation interface{}
}

func (q *Query) Filter(filter interface{}) *Query {
	q.filter = filter
	return q
}

func (q *Query) Sort(sort interface{}) *Query {
	q.sort = sort
	return q
}

func (q *Query) Limit(limit int64) *Query {
	q.limit = limit
	return q
}

func (q *Query) Skip(skip int64) *Query {
	q.skip = skip
	return q
}

//TODO
//func (q *Query) Collation(collation interface{}) *Query {
//    q.collation = collation
//    return q
//}

func (q *Query) Project(proj interface{}) *Query {
	q.proj = proj
	return q
}

func (q *Query) Context(ctx context.Context) *Query {
	q.ctx = ctx
	return q
}

func (q *Query) getFindOptions() []*options.FindOptions {
	opts := make([]*options.FindOptions, 0)
	if q.limit > 0 {
		opts = append(opts, options.Find().SetLimit(q.limit))
	}
	if q.skip > 0 {
		opts = append(opts, options.Find().SetSkip(q.skip))
	}
	if q.sort != nil {
		opts = append(opts, options.Find().SetSort(q.sort))
	}
	if q.proj != nil {
		opts = append(opts, options.Find().SetProjection(q.proj))
	}
	//TODO
	//if q.collation != nil {
	//    opts = append(opts, options.Find().SetCollation(q.collation))
	//}
	return opts
}

func (q *Query) FindOne(data any) error {
	err := q.c.col.FindOne(q.ctx, q.filter).Decode(data)
	return handleError(err)
}

func (q *Query) FindMany(data any) error {
	cursor, err := q.c.col.Find(q.ctx, q.filter, q.getFindOptions()...)
	if err != nil {
		return handleError(err)
	}
	defer cursor.Close(q.ctx)
	return cursor.All(q.ctx, data)
}

func (q *Query) Count() (int64, error) {
	count, err := q.c.col.CountDocuments(q.ctx, q.filter)
	return count, handleError(err)
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
