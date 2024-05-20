package logic

import (
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"lucianaChatServer/conf"
	service "lucianaChatServer/grpc"
	"net"
)

// GrpcService 需要 gRpc 获取数据的函数
func GrpcService() error {
	cerds, err := credentials.NewServerTLSFromFile(conf.CertFile, conf.KeyFile)
	if err != nil {
		return err
	}
	listen, err := net.Listen("tcp", conf.GrpcServerAddress) //local ip and port
	if err != nil {
		return err
	}

	//初始化 gRpc server
	server := grpc.NewServer(grpc.Creds(cerds))

	chat.RegisterGetChatServiceServer(server, &service.GetChat{})
	chat.RegisterGetChatsServiceServer(server, &service.GetChats{})

	if err = server.Serve(listen); err != nil {
		return err
	}

	return nil
}
