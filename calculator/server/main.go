package main

import (
	"log"
	"net"

	pb "github.com/NwinNwin/grpc_learning/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051";

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", addr)

	s := grpc.NewServer()
	// The grpc needs instance for the Greet service, it uses the Server type to implement all the rpc endpoints

	pb.RegisterCalculatorServiceServer(s, &Server{})
	
	if err := s.Serve(lis); err != nil{
		log.Fatalf("Failed to serve: %v", err)
	}
}