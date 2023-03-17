package main

import (
	"context"
	"fmt"

	pricingv1 "github.com/froedevrolijk/grpc-invoicing/proto/pricing/v1"
	"google.golang.org/grpc"
)

type PricingServiceClient struct {
	Client pricingv1.PricingServiceClient
}

func InitPricingServiceClient(url string) PricingServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	client := PricingServiceClient{
		Client: pricingv1.NewPricingServiceClient(cc),
	}

	return client
}

func (s *PricingServiceClient) GetPricing(country string) (*pricingv1.GetPricingResponse, error) {
	req := &pricingv1.GetPricingRequest{
		Order: &pricingv1.Order{
			Country:    pricingv1.Country_NL,
			OrderType:  pricingv1.OrderType_CLAIM,
			PacketSize: pricingv1.PacketSize_SMALL,
		},
	}

	return s.Client.GetPricing(context.Background(), req)

}
