package main

import (
	"log"

	pb "github.com/NwinNwin/grpc_learning/calculator/proto"
)

// this file contains
// implementation of all the rpc endpoints

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error{
	log.Printf("Calulation function was invoked with request %v", in)
	number := in.Number
	var k uint32 = 2;

	for number > 1{
		if number % k == 0{
			stream.Send(&pb.PrimeResponse{
				Result: k})
			number = number / k
		}else{
			k = k + 1
		}
	}
	return nil;
}

// 