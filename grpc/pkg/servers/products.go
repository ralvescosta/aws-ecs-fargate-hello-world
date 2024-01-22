package servers

import (
	"context"

	"github.com/ralvescosta/ec2-hello-world/protos"
	"github.com/ralvescosta/gokit/logging"
)

type productsServer struct {
	logger logging.Logger
	protos.UnimplementedProductsServer
}

func NewProductServer(logger logging.Logger) protos.ProductsServer {
	return &productsServer{
		logger,
		protos.UnimplementedProductsServer{},
	}
}

func (s *productsServer) Create(ctx context.Context, req *protos.CreateProductRequest) (*protos.CreateProductResponse, error) {
	return nil, nil
}

func (s *productsServer) ListProducts(ctx context.Context, req *protos.ListProductsRequest) (*protos.ListProductsResponse, error) {
	return nil, nil
}
