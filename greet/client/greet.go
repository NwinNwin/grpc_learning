package main

import (
	"context"
	"log"

	pb "github.com/NwinNwin/grpc_learning/greet/proto"
)

func doGreet(c pb.GreetServiceClient){
	log.Println("doGreet invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "NwinNwin"})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}
	log.Printf("Greeting: %s", res.Result)
}