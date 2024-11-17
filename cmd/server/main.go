package main

import (
	"context"
	"fmt"
	"github.com/mannanmcc/order/internal/adapter/stock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

	transport "github.com/mannanmcc/order/internal/transport"
	schemas "github.com/mannanmcc/schemas/build/go/rpc/order"
	stockProto "github.com/mannanmcc/schemas/build/go/rpc/stock"
)

func main() {
	port := 50051
	//listen the client request on desire port
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	stClient := buildStockClient()
	resp, _ := stClient.CheckStock(context.Background(), stock.CheckStockRequest{ProductId: 1})
	log.Println("response from stock service, total remaining available", resp.QuantityAvailable)

	//create an instance of gRPC server
	grpcServer := grpc.NewServer()

	//instance of our service implementation
	gRpcServer := transport.New()

	//Register our service implementation with the gRPC server.
	schemas.RegisterOrderServiceServer(grpcServer, gRpcServer)

	//Call Serve() on the server with our port details to do a blocking wait until the process is killed or Stop() is called.
	grpcServer.Serve(lis)
}

func buildStockClient() stock.Client {
	grpcClient, err := grpc.NewClient("0.0.0.0:52051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to create client: %s", err)
	}
	stockClient := stockProto.NewCheckServiceClient(grpcClient)
	return stock.New(stockClient)
}
