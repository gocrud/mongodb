1.安装

```shell
go get github.com/gocrud/mongodb
```

2.使用

```go
// 初始化客户端
client,err := mongodb.New("xxxxx")
if err != nil {
  panic(err)
}
db := client.Database("test")
// 声明集合(Collection)
users := db.Collection("users")
ctx := context.TODO()
// 添加一条数据
user:= map[string]any{
  "name": "test"
}
objectId,err := users.InsertOne(ctx,users)

// 添加多条数据
userList = make([]any,0)
objectIds,err := users.InsertMany(ctx,userList)

// 查询数据
var data User
err := users.Query().Filter(bson.M{"name": "test"}).FindOne(&data)
// 查询多条
var dataArr []User
err := users.Query().Filter(bson.M{"name": "test"}).FindMany(&dataArr)
// 删除数据
err := users.Query().Filter(bson.M{"name": "test"}).DeleteOne()
err := users.Query().Filter(bson.M{"name": "test"}).DeleteMany()
```

