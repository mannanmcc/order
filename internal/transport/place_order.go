package transport

import (
	"context"
	schemas "github.com/mannanmcc/schemas/build/go/rpc/order"
)

func (s Server) PlaceOrder(ctx context.Context, in *schemas.OrderRequest) (*schemas.OrderResponse, error) {
	return &schemas.OrderResponse{OrderId: int32(1)}, nil
}
