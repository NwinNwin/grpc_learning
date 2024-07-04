package main

import (
	"context"
	"log"

	pb "github.com/NwinNwin/grpc_learning/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error){
	/* 
	In here we going to create a Bloitem instance
	and this instance, we are going to try to insert into the database
	if there is an error, we are going to return internal
	- We going to check the ID that MongoDB returned to us
	- If BlogID, we going to return internal, else we will return the BlogID to the client
	*/
	log.Printf("CreateBlog function was invoked with a request: %v", in)

	data := BlogItem{
		AuthorID: in.AuthorId,
		Title: in.Title,
		Content: in.Content,
	}
	// Here we not passing the id because MongoDB will generate the id for us (after we insert it)

	res, err := collection.InsertOne(ctx, data) // Insert to the database

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID) // Get the ID that MongoDB returned to us

	if !ok {
		return nil, status.Errorf(codes.Internal, "Cannot convert to OID")
	}

	return &pb.BlogId{Id: oid.Hex()}, nil
	// Hex is for transforming the object id to a string
}
