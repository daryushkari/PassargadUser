package grpcGateway

import (
	userPb "PassargadUser/api/pb/proto"
	"PassargadUser/entities/requestModel"
	"PassargadUser/usecase"
	"context"
)

type Server struct {
	userPb.UnimplementedUserServer
}

func (s *Server) Create(ctx context.Context, in *userPb.CreateRequest) (*userPb.CreateResponse, error) {
	inp := requestModel.CreateRequest{
		Firstname: in.Firstname,
		Password:  in.Password,
		Lastname:  in.Lastname,
		Email:     in.Email,
		Username:  in.Username,
	}
	err, out := usecase.CreateUser(ctx, inp)
	if err != nil {
		return nil, err
	}
	return &userPb.CreateResponse{
		Message: out.Message,
		Code:    uint32(out.Code),
	}, nil
}
