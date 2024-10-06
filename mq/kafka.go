package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"lucianaChatServer/conf"
	"lucianaChatServer/model"
)

// NewChat 使用kafka发送新聊天信息
func NewChat(uid int64, new *model.ChatInfo) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:  kafka.TCP(conf.KafkaServerAddress),
		Topic: "new_chat",
		//Balancer:               &kafka.Hash{},
		//WriteTimeout:           1 * time.Second,
		//RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", uid)) // key = uid
	value, err := json.Marshal(new)       // value = data
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
func RenameChat(cid int64, name string) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:  kafka.TCP(conf.KafkaServerAddress),
		Topic: "rename",
		//Balancer:               &kafka.Hash{},
		//WriteTimeout:           1 * time.Second,
		//RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", cid)) // key = cid
	value := []byte(name)                 // value = name

	// 发送消息
	err := writer.WriteMessages(
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
func DeleteChat(cid int64) error {
	ctx := context.Background()
	// 创建 Kafka 生产者
	writer := &kafka.Writer{
		Addr:  kafka.TCP(conf.KafkaServerAddress),
		Topic: "delete",
		//Balancer:               &kafka.Hash{},
		//WriteTimeout:           1 * time.Second,
		//RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}

	defer writer.Close()

	// 构造消息
	key := []byte(fmt.Sprintf("%d", cid)) // key = cid
	var value []byte                      // value = nil

	// 发送消息
	err := writer.WriteMessages(
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
