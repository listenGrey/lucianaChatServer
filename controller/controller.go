package controller

import (
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"google.golang.org/grpc"
	"lucianaChatServer/conf"
	service "lucianaChatServer/grpc"
	"net"
)

func ChatService() error {
	/*cerds, err := credentials.NewServerTLSFromFile(conf.CertFile, conf.KeyFile)
	if err != nil {
		return err
	}*/
	listen, err := net.Listen("tcp", conf.GrpcServerAddress) //local ip and port
	if err != nil {
		return err
	}

	//初始化 gRpc server
	server := grpc.NewServer(
	//grpc.Creds(cerds)
	)

	chat.RegisterGetChatListServer(server, &service.ChatList{})
	chat.RegisterGetChatServer(server, &service.GetChat{})
	chat.RegisterNewChatServer(server, &service.NewChat{})
	chat.RegisterRenameChatServer(server, &service.RenameChat{})
	chat.RegisterDeleteChatServer(server, &service.DeleteChat{})

	if err = server.Serve(listen); err != nil {
		return err
	}

	return nil
}
