package main

import (
	"context"
	"fmt"
	"log"
	"lucianaChatServer/controller"
	"lucianaChatServer/dao"
)

func main() {
	var err error
	dao.MongoDB, err = dao.InitClient()
	if err != nil {
		log.Fatalf("mongoDB failed to connect, %s\n", err)
	}
	defer dao.MongoDB.Disconnect(context.TODO())

	fmt.Println("对话服务正在运行")

	err = controller.ChatService()
	if err != nil {
		fmt.Printf("对话服务挂掉了, %s\n", err)
	}
}
