package main

import (
	"google.golang.org/grpc"
	"grpc/internal/pb"
	"grpc/internal/services"
	"net"
)

func main() {

	server := grpc.NewServer()
	pb.RegisterContactServiceServer(server, services.NewContactService())

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}
