package services

import (
	"context"

	"github.com/ralvescosta/ec2-hello-world/protos"
	"github.com/ralvescosta/gokit/logging"
)

type (
	ProductsService interface {
		Create(ctx context.Context, product interface{}) error
		List(context.Context, interface{}) error
	}

	productsService struct {
		logger logging.Logger
		client protos.ProductsClient
	}
)

func NewProductsService(logger logging.Logger, client protos.ProductsClient) ProductsService {
	return &productsService{logger, client}
}

func (s *productsService) Create(ctx context.Context, product interface{}) error {
	return nil
}

func (s *productsService) List(ctx context.Context, product interface{}) error {
	return nil
}
