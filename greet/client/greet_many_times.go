package main

import (
	"context"
	"io"
	"log"

	pb "github.com/NwinNwin/grpc_learning/greet/proto"
)

/*
We will have an infinite loop to call the stream.Recv
from this stream.Recv, we can check if EOF,
if EOF, we will break the loop.
if no error, we will print the response

*/
func doGreetManyTimes(c pb.GreetServiceClient){
	log.Println("doGreetManyTimes invoked")

	req := &pb.GreetRequest{
		FirstName: "NwinNwin",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

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
		log.Printf("Response from GreetManyTimes: %v", msg.Result)
	}
}