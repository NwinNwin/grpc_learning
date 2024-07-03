package main

import (
	"context"
	"log"

	pb "github.com/NwinNwin/grpc_learning/calculator/proto"
)

func calSum(c pb.CalculatorServiceClient){
	log.Println("doGreet invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{First: 1, Second: 2})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}
	log.Printf("Sum from client: %d", res.Result)
}