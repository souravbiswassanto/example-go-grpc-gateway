package internal

import (
	"context"
	"github.com/souravbiswassanto/example-go-grpc-gateway/protogen/golang/orders"
	"log"
)

type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Received an order, adding it to the queue")
	err := o.db.AddOrder(req.GetOrder())
	return &orders.Empty{}, err
}
