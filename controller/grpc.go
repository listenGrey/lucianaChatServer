package controller

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/listenGrey/lucianagRpcPKG/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	service "lucianaChatServer/grpc"
	"net"
)

// GrpcService 需要 gRpc 获取数据的函数
func GrpcService(address string) error {
	serverCert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		return err
	}

	certPool := x509.NewCertPool()
	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		return err
	}
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return errors.New(" Failed to append client CA certificate to pool")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	listen, err := net.Listen("tcp", address) //local ip and port
	if err != nil {
		return err
	}

	//初始化 gRpc server
	server := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))

	chat.RegisterGetChatServiceServer(server, &service.GetChat{})
	chat.RegisterGetChatsServiceServer(server, &service.GetChats{})

	if err = server.Serve(listen); err != nil {
		return err
	}

	return nil
}
