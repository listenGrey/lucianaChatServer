package model

import (
	"encoding/binary"
	"github.com/listenGrey/lucianagRpcPKG/chat"
)

// QA 接收QA
type QA struct {
	Request  string `json:"request" bson:"request"`
	Response string `json:"response" bson:"response"`
}

// ChatInfo 接收 chat info
type ChatInfo struct {
	Cid  int64  `json:"cid"`
	Name string `json:"name"`
}

// Chat 存储在 MongoDB 中的结构
type Chat struct {
	Cid  int64  `bson:"_id"`
	Uid  int64  `bson:"uid"`
	Name string `bson:"name"`
	QAs  []QA   `bson:"qa_s"`
}

func IdMarshal(ori []byte) int64 {
	res := binary.BigEndian.Uint64(ori)
	return int64(res)
}

func ChatsUnmarshal(c *[]Chat) *chat.Chats {
	var res *chat.Chats
	var chats []*chat.Chat

	for _, v := range *c {
		var ch chat.Chat

		ch.Id = v.Cid
		ch.Name = v.Name

		chats = append(chats, &ch)
	}

	res.Chats = chats

	return res
}

func ChatUnmarshal(c *Chat) *chat.Chat {
	var ch *chat.Chat
	var qas []*chat.QA

	for _, v := range c.QAs {
		var qa chat.QA

		qa.Request = v.Request
		qa.Response = v.Response

		qas = append(qas, &qa)
	}

	ch.Id = c.Cid
	ch.Name = c.Name
	ch.Qas = qas

	return ch
}
