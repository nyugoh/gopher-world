package main

import (
	"github.com/nyugoh/gopher-world/RPC/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err.Error())
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)


	if err = grpcServer.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}
