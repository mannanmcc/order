package service

import (
	"context"
	"errors"
	stockClient "github.com/mannanmcc/order/internal/adapter/stock"
	"log"
)

var (
	errFailedToPlaceOrder = errors.New("failed to place order")
	errProductOutOfStock  = errors.New("product is out of stock")
)

type OrderRequest struct {
	ProductId   int32
	Quantity    int32
	TotalAmount int32
}

type OrderResponse struct {
	OrderID int32
}

type StockChecker interface {
	CheckStock(ctx context.Context, req stockClient.CheckStockRequest) (stockClient.CheckStockResponse, error)
}

type Order struct {
	stClient StockChecker
}

func NewOrder(sc StockChecker) Order {
	return Order{
		stClient: sc,
	}
}

func (pl Order) PlaceOrder(ctx context.Context, req OrderRequest) (OrderResponse, error) {
	//validate the request - skipping for simplicity
	//check availability with calling to stock API
	resp, err := pl.stClient.CheckStock(ctx, stockClient.CheckStockRequest{
		ProductId: req.ProductId,
	})

	if err != nil {
		log.Printf("error checking stock %v", err)
		return OrderResponse{}, errFailedToPlaceOrder

	}

	if resp.QuantityAvailable <= req.Quantity {
		return OrderResponse{}, errProductOutOfStock
	}

	//send message to kafka for delivery
	return OrderResponse{OrderID: 1}, nil
}
