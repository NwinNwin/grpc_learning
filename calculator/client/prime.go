package main

import (
	"context"
	"io"
	"log"

	pb "github.com/NwinNwin/grpc_learning/calculator/proto"
)

func calPrime(c pb.CalculatorServiceClient){
	stream, err := c.Prime(context.Background(), &pb.PrimeRequest{Number: 120})

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := stream.Recv()

		// check error if it could be EOF
		if err == io.EOF {break}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		// if here, we don't have error
		log.Printf("Response from Prime: %v", msg.Result)
	}
}