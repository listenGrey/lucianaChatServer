package controller

import (
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"google.golang.org/grpc"
	service "lucianaChatServer/grpc"
	"net"
)

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
