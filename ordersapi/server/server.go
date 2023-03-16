package main

import (
	"context"
	"io"
	"net"
	"os"
	"sync"

	"github.com/froedevrolijk/grpc-invoicing/gateway"
	"github.com/froedevrolijk/grpc-invoicing/insecure"
	ordersv1 "github.com/froedevrolijk/grpc-invoicing/proto/orders/v1"
	"github.com/froedevrolijk/grpc-invoicing/util"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrdersService struct {
	mu     *sync.RWMutex
	orders []*ordersv1.Order
}

func NewOrdersService(orders []*ordersv1.Order) *OrdersService {
	return &OrdersService{
		mu:     &sync.RWMutex{},
		orders: orders,
	}
}

func (s *OrdersService) ListOrders(ctx context.Context, _ *emptypb.Empty) (*ordersv1.ListOrdersResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var orders []*ordersv1.Order
	for _, order := range s.orders {
		orders = append(orders, order)
	}

	return &ordersv1.ListOrdersResponse{Orders: orders}, nil
}

func (s *OrdersService) CreateOrder(ctx context.Context, req *ordersv1.CreateOrderRequest) (*ordersv1.CreateOrderResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	order := &ordersv1.Order{
		Id:     uuid.New().String(),
		Amount: req.Order.Amount,
	}

	s.orders = append(s.orders, order)

	return &ordersv1.CreateOrderResponse{Order: order}, nil
}

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	data, err := util.LoadPbFromCsv("./util/orders.csv")
	if err != nil {
		log.Fatal(err)
	}

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	server := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert))) // TODO: Replace cert

	ordersv1.RegisterOrdersServiceServer(server, NewOrdersService(data))
	reflection.Register(server)

	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(server.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
