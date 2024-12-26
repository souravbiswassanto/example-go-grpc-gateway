package main

import (
	internal "github.com/souravbiswassanto/example-go-grpc-gateway/intenal"
	"github.com/souravbiswassanto/example-go-grpc-gateway/protogen/golang/orders"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	const addr = "0.0.0.0:50051"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	db := internal.NewDB()
	orderService := internal.NewOrderService(db)
	orders.RegisterOrdersServer(server, orderService)
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
