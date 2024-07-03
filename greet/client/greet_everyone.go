package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/NwinNwin/grpc_learning/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient){
	log.Printf("doGreetEveryone invoked")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetEveryone RPC: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "NwinNwin"},
		{FirstName: "NwinNwin1"},
		{FirstName: "NwinNwin2"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs{
			log.Printf("Sending req: %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving response from GreetEveryone: %v", err)
				break
			}
			log.Printf("Received: %v", res.GetResult())
		}
		close(waitc)
	}()

	// make the wait channel wait for the close function call
	// This will wait for the 2 goroutines to finish before exiting the doGreetEveryone function
	<-waitc


	

	

}