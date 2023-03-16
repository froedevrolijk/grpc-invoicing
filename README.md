# grpc-invoicing

### Requirements
Go 1.19+

### Generate OpenAPI spec
`make generate`

### Run server
`go run ordersapi/server/server.go`

### Visit web server
https://0.0.0.0:11000/

### Explore available services
#### List available services
`grpcurl -insecure localhost:10000 list`

#### List available methods for serivce
`grpcurl -insecure localhost:10000 list orders.v1.OrdersService`

#### Describe details for a method
`grpcurl -insecure localhost:10000 describe orders.v1.OrdersService.ListOrders`

#### Describe details for a message
`grpcurl -insecure localhost:10000 describe orders.v1.CreateOrderResponse`

#### Run ListOrders method
`grpcurl -insecure localhost:10000 orders.v1.OrdersService/ListOrders`

#### Run CreateOrder method
```
grpcurl -insecure -d '{
        "order": {
                "amount": "11"
        }
}' localhost:10000 orders.v1.OrdersService/CreateOrder
```