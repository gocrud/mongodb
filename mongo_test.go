package mongodb

import (
	"context"
	"testing"
)

var mdb *Client
var db *Database
var users *Collection

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

func TestCount(t *testing.T) {
	_, err := users.InsertMany(context.TODO(), []any{
		map[string]string{"name": "test1"},
		map[string]string{"name": "test2"},
	})
	if err != nil {
		t.Fatal(err)
	}
	count, err := users.Query().Count()
	if err != nil {
		t.Fatal(err)
	}
	if count == 0 {
		t.Fatal("count error")
	}
	var ls []map[string]string
	err = users.Query().FindMany(&ls)
	if err != nil {
		t.Fatal(err)
	}
	if len(ls) != 2 {
		t.Fatal("find error")
	}
}
