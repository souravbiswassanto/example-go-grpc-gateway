package main

import (
	"fmt"
	"github.com/souravbiswassanto/example-go-grpc-gateway/protogen/golang/orders"
	"github.com/souravbiswassanto/example-go-grpc-gateway/protogen/golang/product"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func main() {
	orderItem := orders.Order{
		OrderId:    10,
		CustomerId: 10,
		IsActive:   true,
		OrderDate:  &date.Date{Year: 2024, Month: 12, Day: 1},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "CocaCola", ProductType: product.ProductType_DRINK},
		},
	}
	b, err := protojson.Marshal(&orderItem)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}
