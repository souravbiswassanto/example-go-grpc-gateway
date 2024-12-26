package internal

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/example-go-grpc-gateway/protogen/golang/orders"
	"log"
)

type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

func NewOrderService(db *DB) *OrderService {
	return &OrderService{
		db: db,
	}
}

func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Received an order, adding it to the queue")
	err := o.db.AddOrder(req.GetOrder())
	return &orders.Empty{}, err
}

func (o *OrderService) GetOrder(_ context.Context, req *orders.PayloadWithOrderID) (*orders.PayloadWithSingleOrder, error) {
	order := o.db.GetOrderById(req.GetOrderId())
	if order == nil {
		return nil, fmt.Errorf("no order found with give order id %v", req.OrderId)
	}
	return &orders.PayloadWithSingleOrder{
		Order: order,
	}, nil
}

func (o *OrderService) UpdateOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Received an update order request")

	o.db.UpdateOrder(req.GetOrder())

	return &orders.Empty{}, nil
}

// RemoveOrder implements the RemoveOrder method of the grpc OrdersServer interface to remove an order
func (o *OrderService) RemoveOrder(_ context.Context, req *orders.PayloadWithOrderID) (*orders.Empty, error) {
	log.Printf("Received a remove order request")

	o.db.RemoveOrder(req.GetOrderId())

	return &orders.Empty{}, nil
}

func (o *OrderService) GetAllOrder(_ context.Context, _ *orders.Empty) (*orders.PayloadWithAllOrders, error) {
	return o.db.GetAllOrder(), nil
}
