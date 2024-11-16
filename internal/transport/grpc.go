package transport

import (
	protos "github.com/mannanmcc/schemas/build/go/rpc/order"
	grpc "google.golang.org/grpc"
)

type Server struct {
	protos.UnimplementedOrderServiceServer
}

func New() *Server {
	return &Server{}
}

// Register register the handler in the GRPC server
func (s *Server) Register(server *grpc.Server) {
	protos.RegisterOrderServiceServer(server, s)
}
