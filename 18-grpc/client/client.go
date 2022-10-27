package main

import (
	"context"
	"fmt"
	"log"

	"grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	service := proto.NewAppServiceClient(clientConn)

	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	ctx := context.Background()
	res, err := service.Add(ctx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Result)
}
