package main

import (
	"code_server/proto"
	"context"
	"fmt"
	"net"
	"sync"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server .
type server struct {
	proto.UnimplementedGreeterServer
	// lock
	mu sync.Mutex
	// 统计出现次数
	count map[string]int
}

// SayHello implementation.
func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count[in.Name]++
	fmt.Println("count: ", s.count[in.Name])
	if s.count[in.Name] > 1 {
		// return one limit error
		st := status.New(codes.ResourceExhausted, "request limit...")
		ds, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{
					{
						Subject:     fmt.Sprintf("name:%s", in.Name),
						Description: "每个name只能调用一次SayHello",
					},
				},
			},
		)
		if err != nil {
			fmt.Printf("st.WithDetails() failed err: %v", err)
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
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
	s := grpc.NewServer()                                                // 创建gRPC服务器
	proto.RegisterGreeterServer(s, &server{count: make(map[string]int)}) // 在gRPC服务端注册服务

	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
