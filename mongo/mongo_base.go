package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var clientMongo *mongo.Client

func GetMongoClient() *mongo.Client {
	return clientMongo
}

// main 中初始化连接
func InitMongoClient(urlMongo, user, password string) error {
	log.Print("Init mongo connection ...")
	client, err := mongo.Connect(TimeoutContext(), options.Client().ApplyURI(urlMongo).SetAuth(options.Credential{
		Username: user,
		Password: password,
	}))
	if err != nil {
		return err
	}
	CloseMongoClient()

	clientMongo = client
	return nil
}

func TimeoutContext() context.Context {
	TimeoutContext, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return TimeoutContext
}

func CloseMongoClient() {
	if clientMongo != nil {
		err := clientMongo.Disconnect(TimeoutContext())
		if err != nil {
			log.Print(err)
		} else {
			log.Print("Close mongo connection ...")
		}
	}
}

// 添加一条记录
func Insert(database string, collection string, doc interface{}) error {
	client := GetMongoClient()
	ctx := TimeoutContext()
	c := client.Database(database).Collection(collection)
	_, err := c.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}

// 添加多条记录
func InsertMany(database string, collection string, docs ...interface{}) error {
	client := GetMongoClient()
	ctx := TimeoutContext()
	c := client.Database(database).Collection(collection)
	_, err := c.InsertMany(ctx, docs)
	if err != nil {
		return err
	}

	return nil
}

// 添加一条记录设置失效时间
func InsertExpire(database string, collection string, doc interface{}, seconds int) error {
	client := GetMongoClient()
	ctx := TimeoutContext()
	c := client.Database(database).Collection(collection)

	_, err := c.InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}

// 统计总数
func Count(database string, collection string, filter interface{}) (int64, error) {
	client := GetMongoClient()
	ctx := TimeoutContext()
	c := client.Database(database).Collection(collection)
	count, err := c.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 分页查询
func FindPage(database string, collection string, skip, limit int64, filter, sort, result interface{}) error {
	client := GetMongoClient()
	ctx := TimeoutContext()
	c := client.Database(database).Collection(collection)
	findOptions := &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
		Sort:  &sort,
	}
	cursor, err := c.Find(ctx, filter, findOptions)
	if err != nil {
		return err
	}

	err = cursor.All(context.TODO(), result)
	if err != nil {
		return err
	}
	return nil
}

// 查找一条记录
func FindOne(database string, collection string, filter, result interface{}) error {
	client := GetMongoClient()
	ctx := TimeoutContext()
	c := client.Database(database).Collection(collection)
	err := c.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

// 更新一条记录
func Update(database string, collection string, filter, update interface{}) error {
	client := GetMongoClient()
	ctx := TimeoutContext()
	c := client.Database(database).Collection(collection)
	_, err := c.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// 设置失效时间 一般不通过代码设置
func ExpireTime(database string, collection string, seconds int) error {
	client := GetMongoClient()
	ctx := TimeoutContext()
	c := client.Database(database).Collection(collection)
	indexModel := mongo.IndexModel{
		Keys:    bsonx.Doc{{"expire_time", bsonx.Int64(1)}},            // 设置TTL索引列
		Options: options.Index().SetExpireAfterSeconds(int32(seconds)), // 设置过期时间
	}
	_, err := c.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	return nil
}

// 更新多条记录
func UpdateMany(database string, collection string, filter, update interface{}) error {
	client := GetMongoClient()
	ctx := TimeoutContext()
	c := client.Database(database).Collection(collection)
	_, err := c.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

// 删除一条记录
func Remove(database string, collection string, filter interface{}) error {
	client := GetMongoClient()
	ctx := TimeoutContext()
	c := client.Database(database).Collection(collection)
	_, err := c.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return err
}

// 删除多条记录
func RemoveMany(database string, collection string, filter interface{}) error {
	client := GetMongoClient()
	ctx := TimeoutContext()
	c := client.Database(database).Collection(collection)
	_, err := c.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	return err
}
