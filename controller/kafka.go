package controller

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"lucianaChatServer/dao"
	"lucianaChatServer/model"
	"time"

	"context"
)

func NewChat(address string) error {
	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "new_chat",
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		var newChat model.Chat
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

func RenameChat(address string) error {
	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "rename",
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		err = dao.RenameChat(ms.Key, ms.Value)
		if err != nil {
			return err
		}
	}
}

func DeleteChat(address string) error {
	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "delete",
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		err = dao.DeleteChat(ms.Key)
		if err != nil {
			return err
		}
	}
}

func SendQA(address string) error {
	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "send_qa",
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
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
