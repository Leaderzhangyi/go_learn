package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-example/route"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runFirst(client pb.RouteGuideClient) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  310235000,
		Longitude: 121437403,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}

func main() {

	// conn, err := grpc.Dial("localhost:10005", grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.NewClient("localhost:10005", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("client cannot dial grpc server")
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)

	runFirst(client)
}
