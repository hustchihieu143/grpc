package gapi

import (
	"context"
	"hustchihieu/grpc_example/pb"
)

func (server *Server) GetAllUser(context.Context, *pb.NoParameterRequest) (*pb.GetAllUserResponse, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method GetAllUser not implemented")
	rsp := &pb.GetAllUserResponse{
		User: "hello",
	}
	return rsp, nil
}
