package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	ctx := context.Background()

	//doRequestResponse(ctx, service)
	//doServerStreaming(ctx, service)
	//doClientStreaming(ctx, service)
	doBiDiStreaming(ctx, service)
}

func doRequestResponse(ctx context.Context, service proto.AppServiceClient) {
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}

	res, err := service.Add(ctx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Result)
}

func doServerStreaming(ctx context.Context, service proto.AppServiceClient) {
	count := 0
	req := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	clientStream, err := service.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		count++
		if count >= 5 {
			clientStream.CloseSend()
			break
		}
		resp, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All the prime nos received")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Prime No received : %d\n", resp.GetNo())
	}
}

func doClientStreaming(ctx context.Context, service proto.AppServiceClient) {
	nos := []int32{3, 1, 4, 2, 5, 8, 6, 9, 7}
	clientStream, err := service.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Sending : %d\n", no)
		req := &proto.AverageRequest{
			No: no,
		}
		err := clientStream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Average Received : ", res.GetAverage())
}

func doBiDiStreaming(ctx context.Context, service proto.AppServiceClient) {
	personNames := []proto.PersonName{
		{FirstName: "Magesh", LastName: "Kuppan"},
		{FirstName: "Suresh", LastName: "Kannan"},
		{FirstName: "Rajesh", LastName: "Pandit"},
		{FirstName: "Ganesh", LastName: "Easwaran"},
		{FirstName: "Ramesh", LastName: "Jayaraman"},
	}
	clientStream, err := service.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	//handle responses
	done := func() <-chan struct{} {
		doneCh := make(chan struct{})
		go func() {
			for {
				res, err := clientStream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalln(err)
				}
				msg := res.GetGreetMessage()
				fmt.Println(msg)
			}
			close(doneCh)
		}()
		return doneCh
	}()

	for _, personName := range personNames {
		fmt.Printf("Sending Person %v\n", personName)
		time.Sleep(500 * time.Millisecond)
		req := &proto.GreetRequest{
			Person: &personName,
		}
		err := clientStream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
	}
	clientStream.CloseSend()
	<-done
}
