package main

import (
	"context"
	"log"
	"time"

	pb "github.com/NwinNwin/grpc_learning/calculator/proto"
)

func calAverage(c pb.CalculatorServiceClient){
	log.Println("doLongGreet invoked")
	
	reqs := []*pb.AverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet RPC: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}
	
	log.Printf("LongGreet Response: %v", res)

}