package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "../helloworld"
	"google.golang.org/grpc"
)

const (
	// 服务器地址
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// 与服务器建立连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn)

	name := defaultName
	// 如果用户有输入参数，则取用户的参数作为name
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用远程过程
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}
