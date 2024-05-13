package controller

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"lucianaChatServer/dao"
	"lucianaChatServer/model"
	"time"

	"context"
)

// NewChat 新建对话
func NewChat(address string) error {
	ctx := context.Background()

	for i := 0; i < 2; i++ {
		go func(groupID string) {
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:        []string{address},
				Topic:          "new_chat",
				CommitInterval: 1 * time.Second,
				GroupID:        groupID,
				StartOffset:    kafka.FirstOffset,
			})

			fmt.Printf("Consumer with GroupID %s is running\n", groupID)

			for {
				ms, err := reader.ReadMessage(ctx)
				if err != nil {
					fmt.Printf("Error reading message for GroupID %s: %s\n", groupID, err)
					continue
				}
				// key = uid
				// value = {cid,name}
				var newChat model.ChatInfo
				err = json.Unmarshal(ms.Value, &newChat)
				if err != nil {
					fmt.Printf("Error unmarshalling message for GroupID %s: %s\n", groupID, err)
					continue
				}

				err = dao.NewChat(ms.Key, &newChat)
				if err != nil {
					fmt.Printf("Error new chat for GroupID %s: %s\n", groupID, err)
					continue
				}

				// 手动提交消息的偏移量
				if err = reader.CommitMessages(ctx, ms); err != nil {
					fmt.Printf("Error committing offset for GroupID %s: %s\n", groupID, err)
				}
			}

		}(fmt.Sprintf("new_chat_chan_%d", i+1))
	}
	select {}
}

// RenameChat 对话重命名
func RenameChat(address string) error {
	ctx := context.Background()

	for i := 0; i < 2; i++ {
		go func(groupID string) {
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:        []string{address},
				Topic:          "rename",
				CommitInterval: 1 * time.Second,
				GroupID:        groupID,
				StartOffset:    kafka.FirstOffset,
			})

			fmt.Printf("Consumer with GroupID %s is running\n", groupID)

			for {
				ms, err := reader.ReadMessage(ctx)
				if err != nil {
					fmt.Printf("Error reading message for GroupID %s: %s\n", groupID, err)
					continue
				}
				// key = cid
				// value = name

				err = dao.RenameChat(ms.Key, ms.Value)
				if err != nil {
					fmt.Printf("Error chat rename for GroupID %s: %s\n", groupID, err)
					continue
				}

				// 手动提交消息的偏移量
				if err = reader.CommitMessages(ctx, ms); err != nil {
					fmt.Printf("Error committing offset for GroupID %s: %s\n", groupID, err)
				}
			}
		}(fmt.Sprintf("rename_chan_%d", i+1))
	}
	select {}
}

// DeleteChat 删除对话
func DeleteChat(address string) error {
	ctx := context.Background()

	for i := 0; i < 2; i++ {
		go func(groupID string) {
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:        []string{address},
				Topic:          "delete",
				CommitInterval: 1 * time.Second,
				GroupID:        groupID,
				StartOffset:    kafka.FirstOffset,
			})

			fmt.Printf("Consumer with GroupID %s is running\n", groupID)

			for {
				ms, err := reader.ReadMessage(ctx)
				if err != nil {
					fmt.Printf("Error reading message for GroupID %s: %s\n", groupID, err)
					continue
				}
				// key = cid
				// value = nil

				err = dao.DeleteChat(ms.Key)
				if err != nil {
					fmt.Printf("Error delete for GroupID %s: %s\n", groupID, err)
					continue
				}

				// 手动提交消息的偏移量
				if err = reader.CommitMessages(ctx, ms); err != nil {
					fmt.Printf("Error committing offset for GroupID %s: %s\n", groupID, err)
				}
			}
		}(fmt.Sprintf("delete_chan_%d", i+1))
	}
	select {}
}

// SendQA 发送QA
func SendQA(address string) error {
	ctx := context.Background()

	for i := 0; i < 2; i++ {
		go func(groupID string) {
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:        []string{address},
				Topic:          "send_qa",
				CommitInterval: 1 * time.Second,
				GroupID:        groupID,
				StartOffset:    kafka.FirstOffset,
			})

			fmt.Printf("Consumer with GroupID %s is running\n", groupID)

			for {
				ms, err := reader.ReadMessage(ctx)
				if err != nil {
					fmt.Printf("Error reading message for GroupID %s: %s\n", groupID, err)
					continue
				}
				// key = cid
				// value = qa
				var qa model.QA
				err = json.Unmarshal(ms.Value, &qa)
				if err != nil {
					fmt.Printf("Error unmarshalling message for GroupID %s: %s\n", groupID, err)
					continue
				}

				err = dao.SendQA(ms.Key, &qa)
				if err != nil {
					fmt.Printf("Error send qa for GroupID %s: %s\n", groupID, err)
					continue
				}

				// 手动提交消息的偏移量
				if err = reader.CommitMessages(ctx, ms); err != nil {
					fmt.Printf("Error committing offset for GroupID %s: %s\n", groupID, err)
				}
			}
		}(fmt.Sprintf("send_qa_chan_%d", i+1))
	}
	select {}
}
