package transport

import (
	"context"
	"github.com/mannanmcc/order/internal/service"
	schemas "github.com/mannanmcc/schemas/build/go/rpc/order"
)

func (s *Server) PlaceOrder(ctx context.Context, in *schemas.OrderRequest) (*schemas.OrderResponse, error) {
	resp, err := s.order.PlaceOrder(ctx, buildRequest(in))
	if err != nil {
		return nil, err
	}

	return &schemas.OrderResponse{OrderId: resp.OrderID}, nil
}

func buildRequest(in *schemas.OrderRequest) service.OrderRequest {
	return service.OrderRequest{
		ProductId:   in.ProductId,
		Quantity:    int32(in.Quantity),
		TotalAmount: in.TotalAmount,
	}
}
