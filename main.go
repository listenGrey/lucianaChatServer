package main

import (
	"fmt"
	"lucianaChatServer/controller"
)

func main() {
	fmt.Println("正在运行")
	errCh1 := make(chan error)
	errCh2 := make(chan error)
	errCh3 := make(chan error)
	errCh4 := make(chan error)
	errCh5 := make(chan error)

	go controller.NewChat(errCh1)
	go controller.RenameChat(errCh2)
	go controller.DeleteChat(errCh3)
	go controller.SendQA(errCh4)
	go controller.GrpcService(errCh5)

	for {
		select {
		case err := <-errCh1:
			fmt.Printf("新建对话挂掉了, %s\n", err)
		case err := <-errCh2:
			fmt.Printf("对话重命名挂掉了, %s\n", err)
		case err := <-errCh3:
			fmt.Printf("删除对话挂掉了, %s\n", err)
		case err := <-errCh4:
			fmt.Printf("发送问题挂掉了, %s\n", err)
		case err := <-errCh5:
			fmt.Printf("gRpc 服务挂掉了, %s \n", err)
		}
	}
}
