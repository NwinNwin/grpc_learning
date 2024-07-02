package main

import (
	"log"
	"net"

	pb "github.com/NwinNwin/grpc_learning/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051";

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", addr)

	s := grpc.NewServer()
	
	if err := s.Serve(lis); err != nil{
		log.Fatalf("Failed to serve: %v", err)
	}
}