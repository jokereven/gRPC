package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mc/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:8000", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	md := metadata.Pairs(
		"token", "jack",
	)

	ctx = metadata.NewOutgoingContext(ctx, md)

	// Get header and trailer（metadata）
	var header, trailer metadata.MD
	resp, err := c.SayHello(
		ctx,
		&proto.HelloRequest{Name: *name},
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		log.Printf("c.SayHello failed, err:%v", err)
		return
	}

	// 拿到响应数据之前可以获取header
	fmt.Printf("header:%v\n", header)

	// 拿到了RPC响应
	log.Printf("resp:%v\n", resp.GetReply())

	// 拿到响应数据后可以获取trailer
	fmt.Printf("trailer:%#v\n", trailer)

	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetReply())
}
