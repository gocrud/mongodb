package mongodb

import (
	"context"
	"testing"
)

var mdb *Client
var db *Database
var users *Collection
var ctx = context.Background()

func TestMain(t *testing.M) {
	var err error
	mdb, err = New("mongodb://root:root@127.0.0.1/?directConnection=true")
	if err != nil {
		panic(err)
	}
	db = mdb.Database("test")
	users = db.Collection("users")
	// 清理数据
	if err := db.Drop(); err != nil {
		panic(err)
	}
	t.Run()
}
