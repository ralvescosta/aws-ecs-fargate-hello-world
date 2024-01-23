package servers

import (
	"context"

	"github.com/ralvescosta/ec2-hello-world/protos"
	"github.com/ralvescosta/ecs-hello-world/grpc/pkg/internal/services"
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
	return nil, nil
}

func (s *productsServer) ListProducts(ctx context.Context, req *protos.ListProductsRequest) (*protos.ListProductsResponse, error) {
	return nil, nil
}
