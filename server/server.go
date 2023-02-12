package main

import (
	"context"
	"log"
	"net"

	ordersv1 "github.com/froedevrolijk/grpc-invoicing/gen/orders/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	ordersv1.UnimplementedOrdersServiceServer
	orders map[string]*ordersv1.Order
}

func NewServer() *server {
	return &server{
		orders: make(map[string]*ordersv1.Order),
	}
}

func (s *server) ListOrders(_ *ordersv1.Empty, stream ordersv1.OrdersService_ListOrdersServer) error {
	for _, order := range s.orders {
		err := stream.Send(&ordersv1.ListOrdersResponse{Order: order})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetOrder(ctx context.Context, req *ordersv1.GetOrderRequest) (*ordersv1.GetOrderResponse, error) {
	order := s.orders[req.Id]
	return &ordersv1.GetOrderResponse{Order: order}, nil
}

func (s *server) CreateOrder(ctx context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {
	req.Order.Id = uuid.New().String()
	s.orders[req.Order.Id] = req.Order
	return &ordersv1.CreateOrderResponse{Order: req.Order}, nil
}

func main() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to create listener: ", err)
	}

	ordersv1.RegisterOrdersServiceServer(s, NewServer())

	log.Println("Starting server on port 8080")

	if err := s.Serve(listener); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
