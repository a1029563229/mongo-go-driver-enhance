package mongoe

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// M: 判断某条数据是否存在于数据库中
func IsExist(ctx context.Context, coll *mongo.Collection, filter interface{}, receiver interface{}) bool {
	result := coll.FindOne(ctx, filter)
	if err := result.Decode(receiver); err != nil {
		fmt.Println("err: ", err)
		return false
	}
	return true
}
