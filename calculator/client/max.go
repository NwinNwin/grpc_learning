package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/NwinNwin/grpc_learning/calculator/proto"
)

func calMax(c pb.CalculatorServiceClient) {
	log.Printf("CalMax invoked")
	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Max RPC: %v", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	waitc := make(chan struct{})

	// send the requests to the server
	go func() {
		for _, req := range reqs {
			log.Printf("sending req: %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// receive the response from the server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving response from Max: %v", err)
				break
			}
			log.Printf("Received: %v", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}