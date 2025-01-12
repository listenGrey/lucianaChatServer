package dao

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"

	"context"
)

var MongoDB *mongo.Client

func init() {
	var err error
	// 设置 MongoDB 连接选项
	clientOptions := options.Client().
		ApplyURI("mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PWD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT") + "/" + os.Getenv("MONGO_DB")).
		SetMaxPoolSize(100).                  // 最大连接池大小
		SetMinPoolSize(10).                   // 最小连接池大小
		SetMaxConnIdleTime(10 * time.Minute). // 最大空闲时间
		SetSocketTimeout(10 * time.Second)    // 连接超时时间

	// 创建一个新的客户端并连接到MongoDB
	MongoDB, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = MongoDB.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
