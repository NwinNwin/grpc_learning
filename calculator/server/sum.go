package main

import (
	"context"
	"log"

	pb "github.com/NwinNwin/grpc_learning/calculator/proto"
)

// this file contains
// implementation of all the rpc endpoints

func (s *Server) Sum(ctx context.Context,in *pb.CalculatorRequest) (*pb.CalculatorResponse, error){
	log.Printf("Calulation function was invoked with request %v", in)
	return &pb.CalculatorResponse{
		Result: in.First + in.Second,
	}, nil
}
// 