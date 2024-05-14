package logic

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"lucianaChatServer/dao"
	"lucianaChatServer/model"
	"time"

	"context"
)

// NewChat 新建对话
func NewChat(address string) error {
	ctx := context.Background()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "new_chat",
		CommitInterval: 1 * time.Second,
		GroupID:        "new_chat",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		// key = uid
		// value = {cid,name}
		var newChat model.ChatInfo
		err = json.Unmarshal(ms.Value, &newChat)
		if err != nil {
			return err
		}

		err = dao.NewChat(ms.Key, &newChat)
		if err != nil {
			return err
		}
	}
}

// RenameChat 对话重命名
func RenameChat(address string) error {
	ctx := context.Background()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "rename",
		CommitInterval: 1 * time.Second,
		GroupID:        "rename",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		// key = cid
		// value = name

		err = dao.RenameChat(ms.Key, ms.Value)
		if err != nil {
			return err
		}
	}
}

// DeleteChat 删除对话
func DeleteChat(address string) error {
	ctx := context.Background()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "delete",
		CommitInterval: 1 * time.Second,
		GroupID:        "delete",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		// key = cid
		// value = nil

		err = dao.DeleteChat(ms.Key)
		if err != nil {
			return err
		}
	}
}

// SendQA 发送QA
func SendQA(address string) error {
	ctx := context.Background()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "send_qa",
		CommitInterval: 1 * time.Second,
		GroupID:        "send_qa",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		// key = cid
		// value = qa
		var qa model.QA
		err = json.Unmarshal(ms.Value, &qa)
		if err != nil {
			return err
		}

		err = dao.SendQA(ms.Key, &qa)
		if err != nil {
			return err
		}
	}
}
