package main

import (
	"context"
	"log"
	"net"

	pb "../helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	// 定义服务端提供服务的端口为50051
	port = ":50051"
)

// 作为实现接口的结构体
type server struct{}

// 实现接口方法SayHello，方法的签名与生成的数据结构定义源码文件stub(*.pb.go)中一致
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello " + in.Name}, nil
}

func main() {
	// 监听50051端口
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建grpc服务端
	s := grpc.NewServer()
	// 注册服务，该方法在生成的stub中定义，由于实现了接口GreetServer的方法SayHello，故方法中可以传递server作为参数
	pb.RegisterGreeterServer(s, &server{})
	// Register refletion service on gRPC server
	reflection.Register(s)
	// 在监听的端口上提供服务
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
