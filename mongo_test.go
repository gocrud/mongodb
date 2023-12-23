package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

var mdb *MongoDB
var db *Database
var users *Collection

func TestMain(t *testing.M) {
	mdb = New("mongodb://root:root@127.0.0.1/?directConnection=true")
	if err := mdb.Connect(); err != nil {
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

// 测试插入一条数据
func TestInsertOne(t *testing.T) {
	objId, err := users.InsertOne(bson.M{"name": "test", "age": 18})
	if err != nil {
		t.Error(err)
	}
	if objId.IsZero() {
		t.Error("insert one failed")
	}
	t.Log(objId)
}

// 测试插入多条数据
func TestInsertMany(t *testing.T) {
	objIds, err := users.InsertMany([]interface{}{
		bson.M{"name": "test1", "age": 18},
		bson.M{"name": "test2", "age": 19},
	})
	if err != nil {
		t.Error(err)
	}
	if len(objIds) != 2 {
		t.Error("insert many failed")
	}
	t.Log(objIds)
}

// 测试查询一条数据
func TestFindOne(t *testing.T) {
	// 添加一条数据
	objId, err := users.InsertOne(bson.M{"name": "test", "age": 18})
	if err != nil {
		t.Error(err)
	}
	if objId.IsZero() {
		t.Error("insert one failed")
	}
	// 查询数据
	var result bson.M
	err = users.Query().Filter(bson.M{"_id": objId}).FindOne(&result)
	if err != nil {
		t.Error(err)
	}
	if result["name"] != "test" {
		t.Error("find one failed")
	}
}

// 测试查询多条数据
func TestFind(t *testing.T) {
	// 添加多条数据
	objIds, err := users.InsertMany([]interface{}{
		bson.M{"name": "test1", "age": 18},
		bson.M{"name": "test2", "age": 19},
	})
	if err != nil {
		t.Error(err)
	}
	if len(objIds) != 2 {
		t.Error("insert many failed")
	}
	// 查询数据
	var result []bson.M
	err = users.Query().Filter(bson.M{"_id": bson.M{"$in": objIds}}).Find(&result)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 2 {
		t.Error("find many failed")
	}
}

// 测试查询数量
func TestCount(t *testing.T) {
	// 添加多条数据
	objIds, err := users.InsertMany([]interface{}{
		bson.M{"name": "test1", "age": 18},
		bson.M{"name": "test2", "age": 19},
	})
	if err != nil {
		t.Error(err)
	}
	if len(objIds) != 2 {
		t.Error("insert many failed")
	}
	// 查询数据
	count, err := users.Query().Filter(bson.M{"_id": bson.M{"$in": objIds}}).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 2 {
		t.Error("count failed")
	}
}

// 测试删除一条数据
func TestDeleteOne(t *testing.T) {
	// 添加一条数据
	objId, err := users.InsertOne(bson.M{"name": "test", "age": 18})
	if err != nil {
		t.Error(err)
	}
	if objId.IsZero() {
		t.Error("insert one failed")
	}
	// 删除数据
	err = users.Query().Filter(bson.M{"_id": objId}).DeleteOne()
	if err != nil {
		t.Error(err)
	}
	// 查询数据
	var result bson.M
	err = users.Query().Filter(bson.M{"_id": objId}).FindOne(&result)
	if err != nil {
		t.Error("delete one failed")
	}
	if result["_id"] != nil {
		t.Error("delete one failed")
	}
}

// 测试删除多条数据
func TestDeleteMany(t *testing.T) {
	// 添加多条数据
	objIds, err := users.InsertMany([]interface{}{
		bson.M{"name": "test1", "age": 18},
		bson.M{"name": "test2", "age": 19},
	})
	if err != nil {
		t.Error(err)
	}
	if len(objIds) != 2 {
		t.Error("insert many failed")
	}
	// 删除数据
	err = users.Query().Filter(bson.M{"_id": bson.M{"$in": objIds}}).DeleteMany()
	if err != nil {
		t.Error(err)
	}
	// 查询数据
	var result []bson.M
	err = users.Query().Filter(bson.M{"_id": bson.M{"$in": objIds}}).Find(&result)
	if err != nil {
		t.Error("delete many failed")
	}
	if len(result) != 0 {
		t.Error("delete many failed")
	}
}

// 测试更新一条数据
func TestUpdateOne(t *testing.T) {
	// 添加一条数据
	objId, err := users.InsertOne(bson.M{"name": "test", "age": 18})
	if err != nil {
		t.Error(err)
	}
	if objId.IsZero() {
		t.Error("insert one failed")
	}
	// 更新数据
	err = users.Query().Filter(bson.M{"_id": objId}).UpdateOne(bson.M{"$set": bson.M{"age": 19}})
	if err != nil {
		t.Error(err)
	}
	// 查询数据
	var result bson.M
	err = users.Query().Filter(bson.M{"_id": objId}).FindOne(&result)
	if err != nil {
		t.Error("update one failed")
	}
	if result["age"].(int32) != 19 {
		t.Error("update one failed")
	}
}

// 测试更新多条数据
func TestUpdateMany(t *testing.T) {
	// 添加多条数据
	objIds, err := users.InsertMany([]interface{}{
		bson.M{"name": "test1", "age": 18},
		bson.M{"name": "test2", "age": 19},
	})
	if err != nil {
		t.Error(err)
	}
	if len(objIds) != 2 {
		t.Error("insert many failed")
	}
	// 更新数据
	err = users.Query().Filter(bson.M{"_id": bson.M{"$in": objIds}}).UpdateMany(bson.M{"$set": bson.M{"age": 20}})
	if err != nil {
		t.Error(err)
	}
	// 查询数据
	var result []bson.M
	err = users.Query().Filter(bson.M{"_id": bson.M{"$in": objIds}}).Find(&result)
	if err != nil {
		t.Error("update many failed")
	}
	for _, v := range result {
		if v["age"].(int32) != 20 {
			t.Error("update many failed")
		}
	}
}

// 测试替换一条数据
func TestReplaceOne(t *testing.T) {
	// 添加一条数据
	objId, err := users.InsertOne(bson.M{"name": "test", "age": 18})
	if err != nil {
		t.Error(err)
	}
	if objId.IsZero() {
		t.Error("insert one failed")
	}
	// 替换数据
	err = users.Query().Filter(bson.M{"_id": objId}).ReplaceOne(bson.M{"name": "test", "age": 19})
	if err != nil {
		t.Error(err)
	}
	// 查询数据
	var result bson.M
	err = users.Query().Filter(bson.M{"_id": objId}).FindOne(&result)
	if err != nil {
		t.Error("replace one failed")
	}
	if result["age"].(int32) != 19 {
		t.Error("replace one failed")
	}
}

// 测试聚合
func TestAggregate(t *testing.T) {
	// 添加多条数据
	objIds, err := users.InsertMany([]interface{}{
		bson.M{"name": "test1", "age": 18},
		bson.M{"name": "test2", "age": 19},
	})
	if err != nil {
		t.Error(err)
	}
	if len(objIds) != 2 {
		t.Error("insert many failed")
	}
	// 聚合
	var result []bson.M
	err = users.Aggregate().Match(bson.M{"_id": bson.M{"$in": objIds}}).Find(&result)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 2 {
		t.Error("aggregate failed")
	}
}

// 测试聚合排序
func TestAggregateSort(t *testing.T) {
	// 添加多条数据
	objIds, err := users.InsertMany([]interface{}{
		bson.M{"name": "test1", "age": 18},
		bson.M{"name": "test2", "age": 19},
	})
	if err != nil {
		t.Error(err)
	}
	if len(objIds) != 2 {
		t.Error("insert many failed")
	}
	// 聚合
	var result []bson.M
	err = users.Aggregate().Match(bson.M{"_id": bson.M{"$in": objIds}}).Sort(bson.M{"age": -1}).Find(&result)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 2 {
		t.Error("aggregate failed")
	}
	if result[0]["age"].(int32) != 19 {
		t.Error("aggregate failed")
	}
}

// 测试聚合分页
func TestAggregatePage(t *testing.T) {
	// 添加多条数据
	objIds, err := users.InsertMany([]interface{}{
		bson.M{"name": "test1", "age": 18},
		bson.M{"name": "test2", "age": 19},
	})
	if err != nil {
		t.Error(err)
	}
	if len(objIds) != 2 {
		t.Error("insert many failed")
	}
	// 聚合
	var result []bson.M
	err = users.Aggregate().Match(bson.M{"_id": bson.M{"$in": objIds}}).Sort(bson.M{"age": -1}).Skip(1).Limit(1).Find(&result)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1 {
		t.Error("aggregate failed")
	}
	if result[0]["age"].(int32) != 18 {
		t.Error("aggregate failed")
	}
}

// 测试聚合分组
func TestAggregateGroup(t *testing.T) {
	// 添加多条数据
	objIds, err := users.InsertMany([]interface{}{
		bson.M{"name": "test1", "age": 18},
		bson.M{"name": "test2", "age": 19},
		bson.M{"name": "test3", "age": 19},
	})
	if err != nil {
		t.Error(err)
	}
	if len(objIds) != 3 {
		t.Error("insert many failed")
	}
	// 聚合
	var result []bson.M
	err = users.Aggregate().Group(bson.M{"_id": "$age", "count": bson.M{"$sum": 1}}).Find(&result)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 2 {
		t.Error("aggregate failed")
	}
	if result[0]["count"].(int32) != 1 {
		t.Error("aggregate failed")
	}
	if result[1]["count"].(int32) != 2 {
		t.Error("aggregate failed")
	}
}
