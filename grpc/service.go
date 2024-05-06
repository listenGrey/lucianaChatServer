package grpc

import (
	"context"
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"lucianaChatServer/dao"
)

type GetChats struct {
	chat.UnimplementedGetChatsServiceServer
}

func (c *GetChats) GetChats(ctx context.Context, uid *chat.ID) (*chat.Chats, error) {
	chats, err := dao.GetChats(uid.GetId())
	if err != nil {
		return nil, err
	}

	return chats, nil
}

type GetChat struct {
	chat.UnimplementedGetChatServiceServer
}

func (c *GetChat) GetChat(ctx context.Context, cid *chat.ID) (*chat.Chat, error) {
	ch, err := dao.GetChat(cid.GetId())
	if err != nil {
		return nil, err
	}

	return ch, nil
}
