package mq

import (
	"lucianaChatServer/model"
	"testing"
)

func TestNewChat(t *testing.T) {
	var uid int64 = 123456
	chat := &model.ChatInfo{
		Cid:  98765,
		Name: "new chat",
	}
	err := NewChat(uid, chat)
	if err != nil {
		t.Error(err)
	}
}

func TestRenameChat(t *testing.T) {
	var cid int64 = 98765
	name := "test rename"
	err := RenameChat(cid, name)
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteChat(t *testing.T) {
	var cid int64 = 98765
	err := DeleteChat(cid)
	if err != nil {
		t.Error(err)
	}
}
