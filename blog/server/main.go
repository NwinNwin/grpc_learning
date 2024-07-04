package main

import (
	"context"
	"log"
	"net"

	pb "github.com/NwinNwin/grpc_learning/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

// Declare Mongo collection to help us with CRUD

/*
SETUP:
- First create a client that will connect to the MongoDB db
- From this client, we will get the db and we going to call that db BlogDB
- From that db, we can access the collection that we will call blog
*/
var collection *mongo.Collection

var addr string = "0.0.0.0:50051";

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	collection = client.Database("blogdb").Collection("blog")



	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", addr)

	s := grpc.NewServer()
	// The grpc needs instance for the Greet service, it uses the Server type to implement all the rpc endpoints

	pb.RegisterBlogServiceServer(s, &Server{})
	
	if err := s.Serve(lis); err != nil{
		log.Fatalf("Failed to serve: %v", err)
	}
}