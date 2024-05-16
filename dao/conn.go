package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lucianaChatServer/conf"

	"context"
)

func MongoDBClient() *mongo.Collection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf.DBAddress))
	if err != nil {
		return nil
	}

	// 选择数据库和集合
	conn := client.Database(conf.Database).Collection(conf.Collection)

	return conn
}
