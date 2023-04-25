package app

import (
	userPb "PassargadUser/api/pb/proto"
	"PassargadUser/config"
	userDelivery "PassargadUser/delivery/grpc/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitApp() {
	cnf, err := config.Get("./config.json")
	if err != nil {
		log.Fatalf("config not available with error: %v", err.Error())
	}

	lis, err := net.Listen("tcp", cnf.ExternalExpose.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := userDelivery.Server{}

	grpcServer := grpc.NewServer()

	userPb.RegisterUserServer(grpcServer, &s)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
