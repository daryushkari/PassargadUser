package main

import (
	userPb "PassargadUser/api/pb/proto"
	userDelivery "PassargadUser/delivery/grpc/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := userDelivery.Server{}

	grpcServer := grpc.NewServer()

	userPb.RegisterUserServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
