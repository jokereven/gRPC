package main

import (
	"context"
	"fmt"
	"net"
	"tls_server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// server .
type server struct {
	proto.UnimplementedGreeterServer
}

// SayHello implementation.
func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	res := "Hello" + " " + in.Name
	return &proto.HelloResponse{Reply: res}, nil
}

func main() {
	// listen 8080 port
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Printf("listen 8000 port failed, err %v", err)
		return
	}

	// 加载证书信息
	creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	if err != nil {
		fmt.Printf("credentials.NewServerTLSFromFile failed, err:%v\n", err)
		return
	}

	s := grpc.NewServer(grpc.Creds(creds))    // 创建gRPC服务器
	proto.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务

	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
