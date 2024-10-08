package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
	"time"
)

func TestMongoDBClient(t *testing.T) {
	client := MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer client.Disconnect(ctx)

	// 选择数据库和集合
	collection := client.Database("luciana").Collection("test")

	// 要插入的数据
	document := bson.D{
		{Key: "name", Value: "John Doe"},
		{Key: "email", Value: "john.doe@example.com"},
		{Key: "age", Value: 30},
	}

	// 插入数据
	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		log.Fatal(err)
	}

	// 查询所有数据
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		t.Error(err)
	}
	defer cursor.Close(ctx)

	// 遍历结果并打印
	for cursor.Next(ctx) {
		var result map[string]interface{}
		err := cursor.Decode(&result)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(result)
	}

	if err := cursor.Err(); err != nil {
		t.Error(err)
	}
}
