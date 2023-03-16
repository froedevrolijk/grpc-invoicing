package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	ordersv1 "github.com/froedevrolijk/grpc-invoicing/proto/orders/v1"
	"github.com/froedevrolijk/grpc-invoicing/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	addr = flag.String("addr", "0.0.0.0", "The address of the server to connect to")
	port = flag.String("port", "10000", "The port to connect to")
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

	createOrderRequest := &ordersv1.CreateOrderRequest{
		Order: &ordersv1.Order{
			Amount: 12,
		},
	}

	fmt.Print("createOrder: ")
	createOrder(client, createOrderRequest)

	fmt.Print("listOrders: ")
	listOrders(client, &emptypb.Empty{})

}

func listOrders(client ordersv1.OrdersServiceClient, empty *emptypb.Empty) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	listedOrders, err := client.ListOrders(ctx, empty)
	util.HandleError(err)
	log.Println(listedOrders)
}

func createOrder(client ordersv1.OrdersServiceClient, req *ordersv1.CreateOrderRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createdOrder, err := client.CreateOrder(ctx, req)
	util.HandleError(err)
	log.Println(createdOrder)
}
