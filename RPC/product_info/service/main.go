package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("App unable to listen:", err.Error())
	}

	grpServer := grpc.NewServer()

	if err = grpServer.Serve(listen); err != nil {
		log.Fatal("gRPC server error:", err.Error())
	}

}
