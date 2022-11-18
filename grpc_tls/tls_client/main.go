package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"
	"tls_client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	creds, err := credentials.NewClientTLSFromFile("certs/server.crt", "jokereven.github.io")
	if err != nil {
		fmt.Printf("credentials.NewClientTLSFromFile failed, err:%v\n", err)
		return
	}

	// 连接到server端，此处禁用安全传输
	conn, err := grpc.Dial(*addr,
		// not safe conn
		// grpc.WithTransportCredentials(insecure.NewCredentials())
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	// 执行RPC调用并打印收到的响应数据
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetReply())
}
