package server

import (
	"sync"

	pricingv1 "github.com/froedevrolijk/grpc-invoicing/gen/pricing/v1"
)

type PricingService struct {
	// pricingv1.UnimplementedPricingServiceServer
	mu      *sync.RWMutex
	pricing []*pricingv1.Pricing
}

func NewPricingService(pricing []*pricingv1.Pricing) *PricingService {
	return &PricingService{pricing: pricing}
}

func (s *PricingService) ListPricing(_ *pricingv1.Empty, stream *pricingv1.PricingService_ListPricingServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, pricing := range s.pricing {
		err := stream.Send(&pricingv1.ListPricingResponse{Pricing: pricing})
		if err != nil {
			return err
		}
	}
	return nil
}
