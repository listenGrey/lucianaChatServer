package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"lucianaChatServer/model"
)

// GetChatList 获取对话列表
func GetChatList(ctx context.Context, uid int64) (*[]model.ChatInfo, error) {
	clo := MongoDB.Database("luciana").Collection("chat_list")
	// 查询指定uid的文档
	var chatList model.ChatList
	filter := bson.M{"_id": uid}
	err := clo.FindOne(ctx, filter).Decode(&chatList)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no chat list found for uid: %d", uid)
		}
		return nil, err
	}

	return &chatList.Chats, nil
}

// GetChat 获取对话
func GetChat(ctx context.Context, cid int64) (*model.Chat, error) {
	clo := MongoDB.Database("luciana").Collection("chat")

	var chats model.Chat
	filter := bson.M{"_id": cid}
	err := clo.FindOne(ctx, filter).Decode(&chats)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no chat found for cid: %d", cid)
		}
		return nil, err
	}

	return &chats, nil
}
