package transport

import (
	"github.com/mannanmcc/order/internal/service"
	schemas "github.com/mannanmcc/schemas/build/go/rpc/order"
)

type Server struct {
	schemas.UnimplementedOrderServiceServer
	order service.Order
}

func New(or service.Order) *Server {
	return &Server{
		order: or,
	}
}
