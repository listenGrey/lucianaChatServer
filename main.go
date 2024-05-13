package main

import (
	"fmt"
	"log"
	"lucianaChatServer/controller"
)

func main() {
	fmt.Println("正在运行")
	if err := controller.NewChat("localhost:9092"); err != nil {
		log.Fatalf("新建对话挂掉了, %s", err)
		return
	}
	if err := controller.RenameChat("localhost:9092"); err != nil {
		log.Fatalf("对话重命名挂掉了, %s", err)
		return
	}
	if err := controller.DeleteChat("localhost:9092"); err != nil {
		log.Fatalf("删除对话挂掉了, %s", err)
		return
	}
	if err := controller.SendQA("localhost:9092"); err != nil {
		log.Fatalf("发送问题挂掉了, %s", err)
		return
	}
	if err := controller.GrpcService("localhost:8964"); err != nil {
		log.Fatalf("gRpc 服务挂掉了, %s", err)
		return
	}
}
