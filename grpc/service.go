package grpc

import (
	"context"
	"github.com/listenGrey/lucianagRpcPKG/chat"
)

type GetChat struct {
	chat.UnimplementedGetChatServiceServer
}

func (c *GetChat) GetChat(ctx context.Context, cid *chat.ID) (*chat.Chat, error) {
	return &chat.Chat{}, nil
}

type GetChats struct {
	chat.UnimplementedGetChatsServiceServer
}

func (c *GetChats) GetChats(ctx context.Context, uid *chat.ID) (*chat.Chats, error) {
	return &chat.Chats{}, nil
}
