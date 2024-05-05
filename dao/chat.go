package dao

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"lucianaChatServer/model"

	"context"
)

func GetChats() {

}

func NewChat(uid []byte, c *model.Chat) error {}

func GetChat() {

}

func RenameChat(cid, name []byte) error {
	client := MongoDBClient("", "")
	if client == nil {
		return errors.New("连接MongoDB失败")
	}

	// 更新数据
	filter := bson.D{{"_id", bson.D{{"$eq", cid}}}}
	update := bson.M{"$set": bson.M{"name": string(name)}}
	_, err := client.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteChat(cid []byte) error {
	client1 := MongoDBClient("", "")
	if client1 == nil {
		return errors.New("连接MongoDB失败")
	}

	filter := bson.D{{"_id", bson.D{{"$eq", cid}}}}
	_, err := client1.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	client2 := MongoDBClient("", "")
	if client2 == nil {
		return errors.New("连接MongoDB失败")
	}

}

func SendQA(cid []byte, qa *model.QA) error {

}
