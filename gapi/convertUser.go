package gapi

import (
	"hustchihieu/grpc_example/pb"
)

func convertUser(user string) *pb.GetAllUserResponse {
	return &pb.GetAllUserResponse{
		User: user,
	}
}
