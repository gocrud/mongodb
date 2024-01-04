package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	cli *mongo.Client
	uri string
}

func New(uri string) (*Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return &Client{
		cli: client,
		uri: uri,
	}, nil
}

func (c *Client) Disconnect() error {
	return c.cli.Disconnect(context.Background())
}

func (c *Client) Transaction(fn func(sessionContext mongo.SessionContext) error) error {
	session, err := c.cli.StartSession()
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

func (c *Client) Database(name string) *Database {
	return &Database{
		name: name,
		db:   c.cli.Database(name),
		mg:   c,
	}
}
