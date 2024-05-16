package controller

import (
	"fmt"
	"lucianaChatServer/logic"
)

func NewChat(errCh chan<- error) {
	fmt.Println("new chat 服务正在运行")
	for {
		// 出现错误则发送到通道
		if err := logic.NewChat(); err != nil {
			errCh <- err
			return
		}
	}
}

func SendQA(errCh chan<- error) {
	fmt.Println("send qa 服务正在运行")
	for {
		// 出现错误则发送到通道
		if err := logic.SendQA(); err != nil {
			errCh <- err
			return
		}
	}
}

func Rename(errCh chan<- error) {
	fmt.Println("rename 服务正在运行")
	for {
		// 出现错误则发送到通道
		if err := logic.RenameChat(); err != nil {
			errCh <- err
			return
		}
	}
}

func Delete(errCh chan<- error) {
	fmt.Println("delete 服务正在运行")
	for {
		// 出现错误则发送到通道
		if err := logic.DeleteChat(); err != nil {
			errCh <- err
			return
		}
	}
}

func GrpcService(errCh chan<- error) {
	fmt.Println("gRpc 服务正在运行")
	for {
		// 出现错误则发送到通道
		if err := logic.GrpcService(); err != nil {
			errCh <- err
			return
		}
	}
}
