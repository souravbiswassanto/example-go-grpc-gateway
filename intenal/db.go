package internal

import (
	"fmt"
	"github.com/souravbiswassanto/example-go-grpc-gateway/protogen/golang/orders"
)

type DB struct {
	collection []*orders.Order
}

func (d *DB) AddOrder(order *orders.Order) error {
	for _, o := range d.collection {
		if o.OrderId == order.OrderId {
			return fmt.Errorf("duplicate order id: %d", order.GetOrderId())
		}
	}
	d.collection = append(d.collection, order)
	return nil
}
