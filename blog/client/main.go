package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/NwinNwin/grpc_learning/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	// Create a connection to the gRPC server
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	// Additional client code goes here...
	c := pb.NewBlogServiceClient(conn)

	// Call function called CreateBlog, takes BlogServiceClient and return id of the Blog created
	readBlog(c, "66860e918d8ac3d5eeb22b37")
}
