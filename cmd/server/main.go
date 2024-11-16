package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	transport "github.com/mannanmcc/order/internal/transport"
	schemas "github.com/mannanmcc/schemas/build/go/rpc/order"
)

func main() {
	port := 50051
	//listen the client request on desire port
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create an instance of gRPC server
	grpcServer := grpc.NewServer()

	//instance of our service implementation
	gRpcServer := transport.New()

	//Register our service implementation with the gRPC server.
	schemas.RegisterOrderServiceServer(grpcServer, gRpcServer)

	//Call Serve() on the server with our port details to do a blocking wait until the process is killed or Stop() is called.
	grpcServer.Serve(lis)
}
