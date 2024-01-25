package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ralvescosta/gokit/guid"
	"github.com/ralvescosta/gokit/logging"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ralvescosta/aws-ecs-fargate-hello-world/grpc/pkg/internal/errors"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/protos"
)

type (
	ProductsService interface {
		Create(ctx context.Context, req *protos.CreateProductRequest) (*protos.CreateProductResponse, error)
		ListProducts(ctx context.Context, req *protos.ListProductsRequest) (*protos.ListProductsResponse, error)
	}

	productsService struct {
		logger   logging.Logger
		products map[string]*protos.Product
	}
)

func NewProductsService(logger logging.Logger) ProductsService {
	return &productsService{
		logger,
		map[string]*protos.Product{},
	}
}

func (s *productsService) Create(ctx context.Context, req *protos.CreateProductRequest) (*protos.CreateProductResponse, error) {
	product := protos.Product{
		Id:        guid.ByteSliceFromStringUUID(uuid.NewString()),
		Name:      req.Name,
		Category:  req.Category,
		Price:     req.Price,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}

	key := s.mapKey(&product)

	if _, ok := s.products[key]; ok {
		s.logger.Warn("product with the same name already created to the same category")
		return nil, errors.NewConflictError(fmt.Sprintf("product with name: %v already created to the same category", req.Name))
	}

	s.products[key] = &product

	return &protos.CreateProductResponse{
		Product: &product,
	}, nil
}

func (s *productsService) ListProducts(ctx context.Context, req *protos.ListProductsRequest) (*protos.ListProductsResponse, error) {
	products := []*protos.Product{}

	for _, v := range s.products {
		if v.Category == req.Category {
			products = append(products, v)
		}
	}

	return &protos.ListProductsResponse{
		Products: products,
	}, nil
}

func (s *productsService) mapKey(req *protos.Product) string {
	return fmt.Sprintf("%v:%v", req.Category, req.Name)
}
