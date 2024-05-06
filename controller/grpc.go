package controller

import (
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"google.golang.org/grpc"
	service "lucianaChatServer/grpc"
	"net"
)

// GrpcService 需要 gRpc 获取数据的函数
func GrpcService(address string) error {
	listen, err := net.Listen("tcp", address) //local ip and port
	if err != nil {
		return err
	}

	//初始化 gRpc server
	server := grpc.NewServer()

	chat.RegisterGetChatServiceServer(server, &service.GetChat{})
	chat.RegisterGetChatsServiceServer(server, &service.GetChats{})

	if err = server.Serve(listen); err != nil {
		return err
	}

	return nil
}
