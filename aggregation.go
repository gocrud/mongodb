package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Aggregation struct {
	c        *Collection
	ctx      context.Context
	pipeline []interface{}
}

func (a *Aggregation) Context(ctx context.Context) *Aggregation {
	a.ctx = ctx
	return a
}

func (a *Aggregation) Clone() *Aggregation {
	aggs := Aggregation{
		c:        a.c,
		ctx:      a.ctx,
		pipeline: a.pipeline,
	}
	return &aggs
}

// Match 过滤
//
//	{ $match: { <query> } }
func (a *Aggregation) Match(filter interface{}) *Aggregation {
	stage := bson.D{{"$match", filter}}
	a.pipeline = append(a.pipeline, stage)
	return a
}

// Group 分组
//
//	{ $group: { _id: <expression>, <field1>: { <accumulator1> : <expression1> }, ... } }
func (a *Aggregation) Group(group interface{}) *Aggregation {
	stage := bson.D{{"$group", group}}
	a.pipeline = append(a.pipeline, stage)
	return a
}

// Project 投影
func (a *Aggregation) Project(project interface{}) *Aggregation {
	stage := bson.D{{"$project", project}}
	a.pipeline = append(a.pipeline, stage)
	return a
}

// Sort 排序
//
//	{ $sort: { <field1>: <sort order>, <field2>: <sort order> ... } }
func (a *Aggregation) Sort(sort interface{}) *Aggregation {
	stage := bson.D{{"$sort", sort}}
	a.pipeline = append(a.pipeline, stage)
	return a
}

// Limit
//
//	{ $limit: <positive integer> }
func (a *Aggregation) Limit(limit int64) *Aggregation {
	stage := bson.D{{"$limit", limit}}
	a.pipeline = append(a.pipeline, stage)
	return a
}

// Skip
//
//	{ $skip: <positive integer> }
func (a *Aggregation) Skip(skip int64) *Aggregation {
	stage := bson.D{{"$skip", skip}}
	a.pipeline = append(a.pipeline, stage)
	return a
}

// Count
//
//	{ $count: <output field name> }
func (a *Aggregation) Count(field string) *Aggregation {
	stage := bson.D{{"$count", field}}
	a.pipeline = append(a.pipeline, stage)
	return a
}

func UnwindOption(path string, includeArrayIndex string, preserveNullAndEmptyArrays bool) bson.M {
	opts := bson.M{
		"preserveNullAndEmptyArrays": preserveNullAndEmptyArrays,
	}
	if path != "" {
		opts["path"] = path
	}
	if includeArrayIndex != "" {
		opts["includeArrayIndex"] = includeArrayIndex
	}
	return opts
}

// Unwind
//
//	{ $unwind: <field path> }
func (a *Aggregation) Unwind(option interface{}) *Aggregation {
	stage := bson.D{{"$unwind", option}}
	a.pipeline = append(a.pipeline, stage)
	return a
}

// ReplaceRoot
//
//	{ $replaceRoot: { newRoot: <expression> } }
func (a *Aggregation) ReplaceRoot(root interface{}) *Aggregation {
	stage := bson.D{{"$replaceRoot", root}}
	a.pipeline = append(a.pipeline, stage)
	return a
}

type LookupOption struct {
	From         string `json:"from" bson:"from"`
	LocalField   string `json:"localField" bson:"localField"`
	ForeignField string `json:"foreignField" bson:"foreignField"`
	As           string `json:"as" bson:"as"`
}

// Lookup
//
//	{ $lookup: { from: <collection to join>, localField: <field from the input documents>, foreignField: <field from the documents of the "from" collection>, as: <output array field> } }
func (a *Aggregation) Lookup(lookup interface{}) *Aggregation {
	stage := bson.D{{"$lookup", lookup}}
	a.pipeline = append(a.pipeline, stage)
	return a
}

func (a *Aggregation) FindOne(result interface{}) error {
	cur, err := a.c.col.Aggregate(a.ctx, a.pipeline)
	if err != nil {
		return handleError(err)
	}
	defer cur.Close(a.ctx)
	if cur.TryNext(a.ctx) {
		err = cur.Decode(result)
	}
	return err
}

func (a *Aggregation) FindMany(result interface{}) error {
	cur, err := a.c.col.Aggregate(a.ctx, a.pipeline)
	if err != nil {
		return handleError(err)
	}
	defer cur.Close(a.ctx)
	return cur.All(a.ctx, result)
}
