package dao

import (
	"errors"
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"go.mongodb.org/mongo-driver/bson"
	"lucianaChatServer/model"

	"context"
)

// GetChats 获取对话列表
func GetChats(uid int64) (*chat.Chats, error) {
	client := MongoDBClient("", "")
	if client == nil {
		return nil, errors.New("连接MongoDB失败")
	}

	filter := bson.D{{"uid", bson.D{{"$eq", uid}}}}
	res, err := client.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var chats []model.Chat
	err = res.Decode(&chats)
	if err != nil {
		return nil, err
	}

	return model.ChatsUnmarshal(&chats), nil
}

// NewChat 新建对话
func NewChat(uid []byte, c *model.ChatInfo) error {
	client := MongoDBClient("", "")
	if client == nil {
		return errors.New("连接MongoDB失败")
	}

	newChat := &model.Chat{
		Cid:  c.Cid,
		Uid:  model.IdMarshal(uid),
		Name: c.Name,
		QAs:  nil,
	}

	newData, err := bson.Marshal(newChat)
	if err != nil {
		return err
	}

	// 插入数据
	_, err = client.InsertOne(context.TODO(), newData)
	if err != nil {
		return err
	}

	return nil
}

// GetChat 获取对话
func GetChat(cid int64) (*chat.Chat, error) {
	client := MongoDBClient("", "")
	if client == nil {
		return nil, errors.New("连接MongoDB失败")
	}

	filter := bson.D{{"cid", bson.D{{"$eq", cid}}}}
	res := client.FindOne(context.TODO(), filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	var ch model.Chat
	err := res.Decode(&ch)
	if err != nil {
		return nil, err
	}

	return model.ChatUnmarshal(&ch), nil
}

// RenameChat 对话重命名
func RenameChat(cid, name []byte) error {
	client := MongoDBClient("", "")
	if client == nil {
		return errors.New("连接MongoDB失败")
	}

	// 更新数据
	filter := bson.D{{"_id", bson.D{{"$eq", model.IdMarshal(cid)}}}}
	update := bson.M{"$set": bson.M{"name": string(name)}}
	_, err := client.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteChat 删除对话
func DeleteChat(cid []byte) error {
	client := MongoDBClient("", "")
	if client == nil {
		return errors.New("连接MongoDB失败")
	}

	filter := bson.D{{"_id", bson.D{{"$eq", model.IdMarshal(cid)}}}}
	_, err := client.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

// SendQA 发送对话
func SendQA(cid []byte, qa *model.QA) error {
	client := MongoDBClient("", "")
	if client == nil {
		return errors.New("连接MongoDB失败")
	}

	filter := bson.D{{"_id", bson.D{{"$eq", model.IdMarshal(cid)}}}}
	oldChat := client.FindOne(context.TODO(), filter)
	if oldChat.Err() != nil {
		return oldChat.Err()
	}

	var newChat model.Chat
	err := oldChat.Decode(&newChat)
	if err != nil {
		return err
	}

	newChat.QAs = append(newChat.QAs, *qa)

	update := bson.M{"$set": bson.M{"qa_s": newChat.QAs}}
	_, err = client.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
