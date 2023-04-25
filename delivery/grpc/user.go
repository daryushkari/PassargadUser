package grpc

import (
	userPb "PassargadUser/api/pb/proto"
	"context"
	"log"
)

type Server struct {
	userPb.UnimplementedUserServer
}

func (s *Server) Create(ctx context.Context, in *userPb.CreateRequest) (*userPb.CreateResponse, error) {
	log.Printf("Receive message body from client: %s", in)
	return &userPb.CreateResponse{}, nil
}
