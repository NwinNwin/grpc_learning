package main

import (
	"context"
	"log"

	pb "github.com/NwinNwin/grpc_learning/greet/proto"
)

// this file contains
// implementation of all the rpc endpoints

func (s *Server) Greet(ctx context.Context,in *pb.GreetRequest) (*pb.GreetResponse, error){
	log.Printf("Greet function was invoked with request %v", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
// 