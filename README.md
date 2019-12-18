# mongo-go-driver 增强版

## 将 struct 转换成 bson

`mongo-go-driver` 自带的增删改查有隐式转换的功能，但是它不能自动忽略空值，并且在指定忽略空值的情况下仍然会插入一些奇怪的数据（比如 time.Time）。

`ToBson` 可以将 `struct` 转换成 `bson` 格式，并且自动忽略空值，可以更好的节约数据库空间，也可以使用 `ToBson` 进行数据的展示。（对 `[]struct` 数据可以使用 `ToBsonList` 进行转换。）

```go
type Prize struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"` 
	Poster      string             `bson:"poster"`                 
	CreatedTime time.Time          `bson:"createdTime"`
	UpdatedTime time.Time          `bson:"updatedTime"`
}

func main() {
  prize := Prize{
    ID:          primitive.NewObjectID(),
    CreatedTime: time.Now(),
  }

  insertData := mongoe.ToBson(prize)
	_, err := *Collection.InsertOne(db.Ctx, insertData)
}
```