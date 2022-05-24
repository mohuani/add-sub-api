package main

import (
	"add-sub-api/mygrpc"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	engine := gin.Default()
	gin.SetMode(gin.DebugMode)
	engine.GET("/createGrpcAddClient", func(c *gin.Context) {
		createGrpcAddClient()
		c.JSON(200, gin.H{
			"message": true,
		})
	})
	engine.GET("/createGrpcSubClient", func(c *gin.Context) {
		createGrpcSubClient()
		c.JSON(200, gin.H{
			"message": true,
		})
	})
	engine.Run(":8080")

	//http.HandleFunc("/createGrpcAddClient", func(writer http.ResponseWriter, request *http.Request) {
	//	createGrpcAddClient()
	//})
	//http.HandleFunc("/createGrpcSubClient", func(writer http.ResponseWriter, request *http.Request) {
	//	createGrpcSubClient()
	//})
	//http.ListenAndServe(":8080", nil)
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
