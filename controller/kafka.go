package controller

import (
	"encoding/json"
	"github.com/segmentio/kafka-go"
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
		var user model.User
		err = json.Unmarshal(ms.Value, &user)
		if err != nil {
			return err
		}

		err = dao.Register(&user)
		if err != nil {
			return err
		}
	}
}

func RenameChat(address string) error {
	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "register",
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		var user model.User
		err = json.Unmarshal(ms.Value, &user)
		if err != nil {
			return err
		}

		err = dao.Register(&user)
		if err != nil {
			return err
		}
	}
}

func DeleteChat(address string) error {
	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "register",
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		var user model.User
		err = json.Unmarshal(ms.Value, &user)
		if err != nil {
			return err
		}

		err = dao.Register(&user)
		if err != nil {
			return err
		}
	}
}

func SendQA(address string) error {
	ctx := context.Background()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{address},
		Topic:          "register",
		CommitInterval: 1 * time.Second,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		ms, err := reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		var user model.User
		err = json.Unmarshal(ms.Value, &user)
		if err != nil {
			return err
		}

		err = dao.Register(&user)
		if err != nil {
			return err
		}
	}
}
