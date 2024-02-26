package servers

import (
	"context"

	"github.com/ralvescosta/aws-ecs-fargate-hello-world/grpc/pkg/internal/errors"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/grpc/pkg/internal/services"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/protos"
	"github.com/ralvescosta/gokit/logging"
)

type productsServer struct {
	logger  logging.Logger
	service services.ProductsService
	protos.UnimplementedProductsServer
}

func NewProductServer(logger logging.Logger, service services.ProductsService) protos.ProductsServer {
	return &productsServer{
		logger,
		service,
		protos.UnimplementedProductsServer{},
	}
}

func (s *productsServer) Create(ctx context.Context, req *protos.CreateProductRequest) (*protos.CreateProductResponse, error) {
	s.logger.Debug("received Create rpc")

	res, err := s.service.Create(ctx, req)
	if err != nil {
		appErr := err.(errors.ApplicationError)
		return nil, appErr.ToGrpc()
	}

	return res, nil
}

func (s *productsServer) ListProducts(ctx context.Context, req *protos.ListProductsRequest) (*protos.ListProductsResponse, error) {
	s.logger.Debug("received ListProducts rpc")

	res, err := s.service.ListProducts(ctx, req)
	if err != nil {
		appErr := err.(errors.ApplicationError)
		return nil, appErr.ToGrpc()
	}

	return res, nil
}
