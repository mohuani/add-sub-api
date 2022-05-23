package main

import (
	"add-sub-api/mygrpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	createGrpcAddClient()
	createGrpcSubClient()
}

func createGrpcAddClient() {

	clientConn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc dail failed: %s\n", err.Error())
	}

	addServiceClient := mygrpc.NewAddServiceClient(clientConn)

	addRequest := mygrpc.AddRequest{
		A: 23,
		B: 45,
	}

	addReply, _ := addServiceClient.Add(context.Background(), &addRequest)

	fmt.Printf("addReply: %s\n", addReply)

}

func createGrpcSubClient() {
	clientConn, err := grpc.Dial(":9998", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("grpc dail failed: %s\n", err.Error())
	}

	subServiceClient := mygrpc.NewSubServiceClient(clientConn)

	subReq := mygrpc.SubReq{
		Number1: 56,
		Number2: 21,
	}

	subResp, err := subServiceClient.Sub(context.Background(), &subReq)

	fmt.Printf("subResp: %s\n", subResp)
}
