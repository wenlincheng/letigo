package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
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
