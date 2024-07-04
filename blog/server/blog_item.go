package main

import (
	pb "github.com/NwinNwin/grpc_learning/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	// this is id type of mongodb, bson will be used to serialize _id and omit if it's empy
	ID primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string `bson:"author_id"`
	Title string `bson:"title"`
	Content string `bson:"content"`
}

// This covert the BlogItem to the pb.Blog onject
func documentToBlog(data *BlogItem) *pb.Blog{
	return &pb.Blog{
		Id: data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title: data.Title,
		Content: data.Content,
	}
}