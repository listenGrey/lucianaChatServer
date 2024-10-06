package grpc

import (
	"context"
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"google.golang.org/grpc/peer"
	"log"
)

type ChatList struct {
	chat.UnimplementedGetChatListServer
}

func (c *ChatList) GetChatList(ctx context.Context, uid *chat.ID) (*chat.ChatList, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("获取对话列表")
	}
	/*chats, err := dao.GetChats(uid.GetId())
	if err != nil {
		return nil, err
	}*/

	return nil, nil
}

type GetChat struct {
	chat.UnimplementedGetChatServer
}

func (c *GetChat) GetChat(ctx context.Context, cid *chat.ID) (*chat.Chat, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("获取对话")
	}
	/*ch, err := dao.GetChat(cid.GetId())
	if err != nil {
		return nil, err
	}*/

	return nil, nil
}

type NewChat struct {
	chat.UnimplementedNewChatServer
}

func (c *NewChat) NewChat(ctx context.Context, uc *chat.UserChat) (*chat.Null, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("新建对话")
	}
	return nil, nil
}

type RenameChat struct {
	chat.UnimplementedRenameChatServer
}

func (c *RenameChat) RenameChat(ctx context.Context, ch *chat.Chat) (*chat.Null, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("对话重命名")
	}
	return nil, nil
}

type DeleteChat struct {
	chat.UnimplementedDeleteChatServer
}

func (c *DeleteChat) DeleteChat(ctx context.Context, cid *chat.ID) (*chat.Null, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("删除对话")
	}
	return nil, nil
}
