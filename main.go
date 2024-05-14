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

	go controller.New(errCh1)
	go controller.Modify(errCh2)
	go controller.GrpcService(errCh3)

	for {
		select {
		case err := <-errCh1:
			fmt.Printf("新建对话挂掉了, %s\n", err)
			fmt.Printf("发送问题挂掉了, %s\n", err)
		case err := <-errCh2:
			fmt.Printf("对话重命名挂掉了, %s\n", err)
			fmt.Printf("删除对话挂掉了, %s\n", err)
		case err := <-errCh3:
			fmt.Printf("gRpc 服务挂掉了, %s \n", err)
		}
	}
}
