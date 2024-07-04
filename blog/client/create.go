package main

import (
	"context"
	"log"

	pb "github.com/NwinNwin/grpc_learning/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--- Create Blog was invoked ---")
	blog := &pb.Blog{
		AuthorId: "Nwin",
		Title: "My First Blog",
		Content: "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Error while creating blog: %v", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id

}