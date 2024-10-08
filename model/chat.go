package model

import (
	"github.com/listenGrey/lucianagRpcPKG/chat"
)

type ChatList struct {
	Uid   int64      `bson:"_id"`
	Chats []ChatInfo `bson:"chats"`
}

// ChatInfo 接收 chat info
type ChatInfo struct {
	Cid  int64  `bson:"cid"`
	Name string `bson:"name"`
}

// Chat 存储在 MongoDB 中的结构
type Chat struct {
	Cid  int64  `bson:"_id"`
	Uid  int64  `bson:"uid"`
	Name string `bson:"name"`
	QAs  []QA   `bson:"qas"`
}

// QA 接收QA
type QA struct {
	Request  string `bson:"request"`
	Response string `bson:"response"`
}

func ChatListsUnmarshal(uid int64, c *[]ChatInfo) *chat.ChatList {
	var res *chat.ChatList
	var chats []*chat.ChatInfo

	res.Uid = uid

	for _, v := range *c {
		var ch *chat.ChatInfo

		ch.Cid = v.Cid
		ch.Name = v.Name

		chats = append(chats, ch)
	}
	res.Chats = chats
	return res
}

func ChatUnmarshal(c *Chat) *chat.Chat {
	var ch *chat.Chat
	var qas []*chat.QA

	for _, v := range c.QAs {
		var qa *chat.QA

		qa.Request = v.Request
		qa.Response = v.Response

		qas = append(qas, qa)
	}

	ch.Cid = c.Cid
	ch.Name = c.Name
	ch.Qas = qas

	return ch
}

func ChatMarshal(c *chat.Chat) *Chat {
	var ch *Chat
	var qas []QA

	for _, v := range c.Qas {
		var qa QA

		qa.Request = v.Request
		qa.Response = v.Response

		qas = append(qas, qa)
	}

	ch.Cid = c.Cid
	ch.Name = c.Name
	ch.QAs = qas

	return ch

}

func ChatInfoMarshal(c *chat.ChatInfo) *ChatInfo {
	return &ChatInfo{
		Cid:  c.GetCid(),
		Name: c.GetName(),
	}
}
