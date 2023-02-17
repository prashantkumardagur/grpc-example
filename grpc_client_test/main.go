package main

import (
	"context"
	"log"

	"grpc_client_example/news"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// Register our service
	client := news.NewNewsServiceClient(conn)

	//==================================================================================================

	// // Call the GetAllṇews method
	// response, err := c.GetAllNews(context.Background(), &news.Empty{})
	// if err != nil {
	// 	log.Fatalf("Error when getting all news")
	// }
	// log.Printf("Response from server: %v", response)

	// Call the GetNewsById method
	response, err := client.GetNewsById(context.Background(), &news.News{Id: "1"})
	if err != nil {
		log.Fatalf("Error when getting news by id")
	}
	log.Printf("Response from server: %v", response)

}
