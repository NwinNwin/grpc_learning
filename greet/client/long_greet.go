package main

import (
	"context"
	"log"
	"time"

	pb "github.com/NwinNwin/grpc_learning/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient){
	log.Println("doLongGreet invoked")
	
	reqs := []*pb.GreetRequest{
	{FirstName: "NwinNwin"},
	{FirstName: "NwinNwin1"},
	{FirstName: "NwinNwin2"},
	}

	stream, err := c.LongGreet(context.Background())

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