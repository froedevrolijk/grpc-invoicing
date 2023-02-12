package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/froedevrolijk/grpc-invoicing/common"
	ordersv1 "github.com/froedevrolijk/grpc-invoicing/gen/orders/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost", "The address of the server to connect to")
	port = flag.String("port", "8080", "The port to connect to")
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	conn, err := grpc.DialContext(ctx, net.JoinHostPort(*addr, *port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Failed to connect to server: ", err)
	}

	client := ordersv1.NewOrdersServiceClient(conn)

	fmt.Println("listOrders")
	listOrders(client, &ordersv1.Empty{})

	createOrderRequest := &ordersv1.CreateOrderRequest{
		Order: &ordersv1.Order{
			Amount: 12,
		},
	}

	fmt.Println("createOrder")
	createOrder(client, createOrderRequest)
}

func listOrders(client ordersv1.OrdersServiceClient, empty *ordersv1.Empty) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ListOrders(ctx, empty)
	common.HandleError(err)
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			break
		}
		common.HandleError(err)
		log.Printf("Order: id: %v, amount: %v", order.Order.Id, order.Order.Amount)
	}
}

func createOrder(client ordersv1.OrdersServiceClient, req *ordersv1.CreateOrderRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createdOrder, err := client.CreateOrder(ctx, req)
	common.HandleError(err)
	log.Println(createdOrder)
}
