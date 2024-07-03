package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/NwinNwin/grpc_learning/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with a streaming request")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Printf("Error while reading client stream: %v", err)
		}

		res += fmt.Sprintf("Hello %s! ", req.FirstName)
	}
}