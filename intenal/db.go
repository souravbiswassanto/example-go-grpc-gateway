package internal

import (
	"fmt"
	"github.com/souravbiswassanto/example-go-grpc-gateway/protogen/golang/orders"
)

type DB struct {
	collection []*orders.Order
}

func NewDB() *DB {
	return &DB{
		collection: make([]*orders.Order, 0),
	}
}

func (db *DB) AddOrder(order *orders.Order) error {
	if order == nil {
		return fmt.Errorf("order is empty")
	}
	for _, o := range db.collection {
		if o.OrderId == order.OrderId {
			return fmt.Errorf("duplicate order id: %d", order.GetOrderId())
		}
	}
	db.collection = append(db.collection, order)
	return nil
}

func (db *DB) GetOrderById(orderId uint64) *orders.Order {
	for _, o := range db.collection {
		if o.OrderId == orderId {
			return o
		}
	}
	return nil
}

// GetOrdersByIDs returns all orders pertaining to the given order ids
func (d *DB) GetOrdersByIDs(orderIDs []uint64) []*orders.Order {
	filtered := make([]*orders.Order, 0)

	for _, idx := range orderIDs {
		for _, order := range d.collection {
			if order.OrderId == idx {
				filtered = append(filtered, order)
				break
			}
		}
	}

	return filtered
}

func (d *DB) GetAllOrder() *orders.PayloadWithAllOrders {
	return &orders.PayloadWithAllOrders{
		Order: d.collection,
	}
}

// UpdateOrder updates an order in place
func (d *DB) UpdateOrder(order *orders.Order) {
	for i, o := range d.collection {
		if o.OrderId == order.OrderId {
			d.collection[i] = order
			return
		}
	}
}

// RemoveOrder removes an order from the orders collection
func (d *DB) RemoveOrder(orderID uint64) {
	filtered := make([]*orders.Order, 0, len(d.collection)-1)
	for i := range d.collection {
		if d.collection[i].OrderId != orderID {
			filtered = append(filtered, d.collection[i])
		}
	}
	d.collection = filtered
}
