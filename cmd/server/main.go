package main

import (
	"fmt"
	"github.com/mannanmcc/order/internal/adapter/stock"
	"github.com/mannanmcc/order/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"

	"github.com/mannanmcc/order/internal/config"
	"github.com/mannanmcc/order/internal/transport"
	schemas "github.com/mannanmcc/schemas/build/go/rpc/order"
	stockProto "github.com/mannanmcc/schemas/build/go/rpc/stock"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	stClient, err := buildStockClient(cfg)
	if err != nil {
		log.Fatalf("failed to create connection to stock API: %v", err)
		return
	}

	order := service.NewOrder(stClient)
	//create an instance of gRPC server
	grpcServer := grpc.NewServer()

	//instance of our service implementation
	gRpcServer := transport.New(order)

	//Register our service implementation with the gRPC server.
	schemas.RegisterOrderServiceServer(grpcServer, gRpcServer)

	port := 50051
	//Call Serve() on the server with our port details to do a blocking wait until the process is killed or Stop() is called.
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatal("Unable to listen on port:", "50051")
		os.Exit(1)
	}

	//start the grpc server
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("failed to serve grpc server: %v", err)
	}
}

func buildStockClient(cfg config.Config) (*stock.Client, error) {
	grpcClient, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.StockHostName, cfg.StockHostPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Printf("connecting stock server on :%s and :%s", cfg.StockHostName, cfg.StockHostPort)
	if err != nil {
		log.Printf("failed to create client: %s", err)
		return nil, err
	}
	stockClient := stockProto.NewCheckServiceClient(grpcClient)
	return stock.New(stockClient), nil
}
