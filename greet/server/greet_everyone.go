package main

import (
	"io"
	"log"

	pb "github.com/NwinNwin/grpc_learning/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked with a streaming request")
	
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		
		if err != nil {
			log.Printf("Error while reading client stream: %v",err )}
		
		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,})
		
		if err != nil {
			log.Printf("Error while sending data to client: %v",err )}
	}
}
