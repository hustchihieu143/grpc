package main

import (
	"hustchihieu/grpc_example/gapi"
	"hustchihieu/grpc_example/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	runGrpcServer()
}

func runGrpcServer() {
	server, err := gapi.NewServer()

	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:9096")
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Println("start gRPC server at: ", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}
