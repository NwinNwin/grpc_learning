package main

import (
	"io"
	"log"

	pb "github.com/NwinNwin/grpc_learning/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Avg function was invoked with a streaming request")

	var res float32 = 0
	var count float32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: res / count,
			})
		}

		if err != nil {
			log.Printf("Error while reading client stream: %v", err)
		}

		res += req.Number
		count++
	}
}