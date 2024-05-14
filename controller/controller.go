package controller

import (
	"fmt"
	"lucianaChatServer/conf"
	"lucianaChatServer/logic"
)

func New(errCh chan<- error) {
	fmt.Println("new chat 服务正在运行")
	fmt.Println("send qa 服务正在运行")
	for {
		// 出现错误则发送到通道
		if err := logic.NewChat(conf.KafkaServerAddress); err != nil {
			errCh <- err
			return
		}
		// 出现错误则发送到通道
		if err := logic.SendQA(conf.KafkaServerAddress); err != nil {
			errCh <- err
			return
		}
	}
}

func Modify(errCh chan<- error) {
	fmt.Println("rename 服务正在运行")
	fmt.Println("delete 服务正在运行")
	for {
		// 出现错误则发送到通道
		if err := logic.RenameChat(conf.KafkaServerAddress); err != nil {
			errCh <- err
			return
		}
		// 出现错误则发送到通道
		if err := logic.DeleteChat(conf.KafkaServerAddress); err != nil {
			errCh <- err
			return
		}
	}
}

func GrpcService(errCh chan<- error) {
	fmt.Println("gRpc 服务正在运行")
	for {
		// 出现错误则发送到通道
		if err := logic.GrpcService(conf.GrpcServerAddress); err != nil {
			errCh <- err
			return
		}
	}
}
