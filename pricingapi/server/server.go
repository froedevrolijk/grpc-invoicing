package server

import (
	"sync"

	"github.com/froedevrolijk/grpc-invoicing/pricingapi/db"
	"github.com/froedevrolijk/grpc-invoicing/pricingapi/models"
	"golang.org/x/net/context"

	pricingv1 "github.com/froedevrolijk/grpc-invoicing/proto/pricing/v1"
)

type PricingService struct {
	mu *sync.RWMutex
	H  db.Handler
}

func (s *PricingService) GetPricing(ctx context.Context, req *pricingv1.Order) (*pricingv1.GetPricingResponse, error) {
	var pricing models.Pricing

}
