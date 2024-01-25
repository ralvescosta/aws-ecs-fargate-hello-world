package services

import (
	"context"

	"github.com/ralvescosta/aws-ecs-fargate-hello-world/protos"
	"github.com/ralvescosta/gokit/logging"
)

type (
	ProductsService interface {
		Create(context.Context, *protos.CreateProductRequest) (*protos.CreateProductResponse, error)
		List(context.Context, *protos.ListProductsRequest) (*protos.ListProductsResponse, error)
	}

	productsService struct {
		logger logging.Logger
		client protos.ProductsClient
	}
)

func NewProductsService(logger logging.Logger, client protos.ProductsClient) ProductsService {
	return &productsService{logger, client}
}

func (s *productsService) Create(ctx context.Context, req *protos.CreateProductRequest) (*protos.CreateProductResponse, error) {
	return s.client.Create(ctx, req)
}

func (s *productsService) List(ctx context.Context, req *protos.ListProductsRequest) (*protos.ListProductsResponse, error) {
	return s.client.ListProducts(ctx, req)
}
