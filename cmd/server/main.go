package main

import (
	"fmt"
	"github.com/mannanmcc/order/internal/adapter/stock"
	"github.com/mannanmcc/order/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

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
	//listen the client request on desire port
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.StockHostName, cfg.StockHostPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	stClient, err := buildStockClient()
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

	//Call Serve() on the server with our port details to do a blocking wait until the process is killed or Stop() is called.
	grpcServer.Serve(lis)
}

func buildStockClient() (stock.Client, error) {
	grpcClient, err := grpc.NewClient("0.0.0.0:52051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to create client: %s", err)
		return stock.Client{}, err
	}
	stockClient := stockProto.NewCheckServiceClient(grpcClient)
	return stock.New(stockClient), nil
}
