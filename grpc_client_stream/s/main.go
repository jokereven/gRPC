package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"s/proto"

	"google.golang.org/grpc"
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

// LotsOfReplies 返回使用多种语言打招呼
func (s *server) LotsOfReplies(in *proto.HelloRequest, stream proto.Greeter_LotsOfRepliesServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}

	for _, word := range words {
		data := &proto.HelloResponse{
			Reply: word + in.GetName(),
		}
		// 使用Send方法返回多个数据
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

// LotsOfGreetings 接收流式数据
func (s *server) LotsOfGreetings(stream proto.Greeter_LotsOfGreetingsServer) error {
	reply := "你好："
	for {
		// 接收客户端发来的流式数据
		res, err := stream.Recv()
		if err == io.EOF {
			// 最终统一回复
			return stream.SendAndClose(&proto.HelloResponse{
				Reply: reply,
			})
		}
		if err != nil {
			return err
		}
		reply += res.GetName()
	}
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
