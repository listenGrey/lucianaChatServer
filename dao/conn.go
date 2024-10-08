package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"lucianaChatServer/conf"
	"time"

	"context"
)

var MongoDB *mongo.Client

func InitClient() (*mongo.Client, error) {
	// 设置连接选项
	clientOptions := options.Client().ApplyURI(conf.DBAddress).SetConnectTimeout(10 * time.Second).SetSocketTimeout(1 * time.Second)

	// 创建一个新的客户端并连接到MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	return client, nil
}
