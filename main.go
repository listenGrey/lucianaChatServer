package main

import (
	"fmt"
	"lucianaChatServer/controller"
)

func main() {
	fmt.Println("对话服务正在运行")
	err := controller.ChatService()
	if err != nil {
		fmt.Printf("对话服务挂掉了, %s\n", err)
	}
}
