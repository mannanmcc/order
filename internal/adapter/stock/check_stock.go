package stock

import (
	"context"
	"errors"
	"github.com/mannanmcc/schemas/build/go/rpc/stock"
	"google.golang.org/grpc"
)

var (
	ErrorFailedToCommunicateStockService = errors.New("failed to communcate stock service")
)

type Stock interface {
	CheckStock(ctx context.Context, in *stock.CheckStockRequest, opts ...grpc.CallOption) (*stock.CheckStockResponse, error)
}

type Client struct {
	stockClient Stock
}

func New(cl Stock) Client {
	return Client{stockClient: cl}
}

type CheckStockRequest struct {
	ProductId int32
}

type CheckStockResponse struct {
	ProductID         int32
	QuantityAvailable int32
}

func (c Client) CheckStock(ctx context.Context, req CheckStockRequest) (CheckStockResponse, error) {
	resp, err := c.stockClient.CheckStock(ctx, &stock.CheckStockRequest{ProductId: req.ProductId})

	if err != nil {
		return CheckStockResponse{}, ErrorFailedToCommunicateStockService
	}

	return CheckStockResponse{
		QuantityAvailable: resp.QuantityAvailable,
	}, nil
}
