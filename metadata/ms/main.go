package main

import (
	"context"
	"fmt"
	"ms/proto"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// server .
type server struct {
	proto.UnimplementedGreeterServer
}

// SayHello implementation.
func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	defer func() {
		trailer := metadata.Pairs(
			"timestamp", strconv.Itoa(int(time.Now().Unix())),
		)
		grpc.SetTrailer(ctx, trailer)
	}()

	// check the metadata have or no token
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("md:%#v ok:%#v\n", md, ok)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "invalid request")
	}
	vl := md.Get("token")
	if len(vl) < 1 || vl[0] != "jack" {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	header := metadata.New(map[string]string{"location": "Hangzhou"})
	grpc.SendHeader(ctx, header)

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
	s := grpc.NewServer()                     // 创建gRPC服务器
	proto.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务

	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
