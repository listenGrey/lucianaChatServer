package grpc

import (
	"context"
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"google.golang.org/grpc/peer"
	"log"
	"lucianaChatServer/dao"
	"lucianaChatServer/model"
	"lucianaChatServer/mq"
)

type ChatList struct {
	chat.UnimplementedGetChatListServer
}

func (c *ChatList) GetChatList(ctx context.Context, uid *chat.ID) (*chat.ChatList, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("获取对话列表")
	}
	chatInfos, err := dao.GetChatList(ctx, uid.GetId())
	if err != nil {
		return nil, err
	}

	return model.ChatListsUnmarshal(uid.GetId(), chatInfos), nil
}

type GetChat struct {
	chat.UnimplementedGetChatServer
}

func (c *GetChat) GetChat(ctx context.Context, cid *chat.ID) (*chat.Chat, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("获取对话")
	}
	chats, err := dao.GetChat(ctx, cid.GetId())
	if err != nil {
		return nil, err
	}

	return model.ChatUnmarshal(chats), nil
}

type NewChat struct {
	chat.UnimplementedNewChatServer
}

func (c *NewChat) NewChat(ctx context.Context, ch *chat.Chat) (*chat.Null, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("新建对话")
	}

	newChat := model.ChatMarshal(ch)
	err := mq.NewChat(newChat)
	if err != nil {
		return &chat.Null{}, err
	}

	return &chat.Null{}, nil
}

type RenameChat struct {
	chat.UnimplementedRenameChatServer
}

func (c *RenameChat) RenameChat(ctx context.Context, ch *chat.ChatInfo) (*chat.Null, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("对话重命名")
	}

	chats := model.ChatInfoMarshal(ch)
	err := mq.RenameChat(chats)
	if err != nil {
		return &chat.Null{}, err
	}

	return &chat.Null{}, nil
}

type DeleteChat struct {
	chat.UnimplementedDeleteChatServer
}

func (c *DeleteChat) DeleteChat(ctx context.Context, cid *chat.ID) (*chat.Null, error) {
	_, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("删除对话")
	}

	err := mq.DeleteChat(cid.GetId())
	if err != nil {
		return &chat.Null{}, err
	}
	return &chat.Null{}, nil
}
