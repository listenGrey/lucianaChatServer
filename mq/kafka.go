package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
	"lucianaChatServer/model"
)

// NewChat 使用kafka发送新聊天信息
func NewChat(newChat *model.Chat) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:  kafka.TCP(os.Getenv("KAFKA_ADDR") + ":" + os.Getenv("KAFKA_PORT")),
		Topic: "new_chat",
		//Balancer:               &kafka.Hash{},
		//WriteTimeout:           1 * time.Second,
		//RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", newChat.Uid)) // key = uid
	value, err := json.Marshal(newChat)           // value = data
	if err != nil {
		return err
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// RenameChat 使用kafka发送修改聊天名
func RenameChat(chat *model.Chat) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:  kafka.TCP(os.Getenv("KAFKA_ADDR") + ":" + os.Getenv("KAFKA_PORT")),
		Topic: "rename",
		//Balancer:               &kafka.Hash{},
		//WriteTimeout:           1 * time.Second,
		//RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", chat.Uid)) // key = cid
	value, err := json.Marshal(chat)           // value = data
	if err != nil {
		return err
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteChat 使用kafka发送删除聊天
func DeleteChat(chat *model.Chat) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:  kafka.TCP(os.Getenv("KAFKA_ADDR") + ":" + os.Getenv("KAFKA_PORT")),
		Topic: "delete",
		//Balancer:               &kafka.Hash{},
		//WriteTimeout:           1 * time.Second,
		//RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", chat.Uid)) // key = cid
	value, err := json.Marshal(chat)           // value = data
	if err != nil {
		return err
	}

	// 发送消息
	err = writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
