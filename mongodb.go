package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	cli *mongo.Client
	uri string
}

func New(uri string) *MongoDB {
	return &MongoDB{uri: uri}
}

func (m *MongoDB) Connect() error {
	clientOptions := options.Client().ApplyURI(m.uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	m.cli = client
	return nil
}

func (m *MongoDB) Disconnect() error {
	return m.cli.Disconnect(context.Background())
}

func (m *MongoDB) Transaction(fn func(sessionContext mongo.SessionContext) error) error {
	session, err := m.cli.StartSession()
	if err != nil {
		return err
	}
	defer func() {
		session.EndSession(context.Background())
	}()

	var f = func(sessionContext mongo.SessionContext) (interface{}, error) {
		err := fn(sessionContext)
		return nil, err
	}

	_, err = session.WithTransaction(context.Background(), f)
	return err
}

func (m *MongoDB) Database(name string) *Database {
	return &Database{
		name: name,
		db:   m.cli.Database(name),
		mg:   m,
	}
}
