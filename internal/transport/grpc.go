package transport

import (
	schemas "github.com/mannanmcc/schemas/build/go/rpc/order"
)

type Server struct {
	schemas.UnimplementedOrderServiceServer
}

func New() *Server {
	return &Server{}
}
