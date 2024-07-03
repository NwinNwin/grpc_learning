package main

import (
	"io"
	"log"

	pb "github.com/NwinNwin/grpc_learning/calculator/proto"
)


func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked with a streaming request")

	var curMax uint32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error while reading client stream: %v", err)
		}

		curMax = max(curMax, req.Number)
		err = stream.Send(&pb.MaxResponse{
			Result: curMax,
		})

		if err != nil {
			log.Printf("Error while sending data to client: %v", err)
		}
	}
}
