package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ralvescosta/aws-ecs-fargate-hello-world/api/pkg/internal/services"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/api/pkg/views"
	"github.com/ralvescosta/aws-ecs-fargate-hello-world/protos"
	"github.com/ralvescosta/gokit/httpw/server"
	"github.com/ralvescosta/gokit/logging"
	"go.uber.org/zap"
)

type (
	HTTPHandlers interface {
		Install(router server.HTTPServer)
	}

	productsHandler struct {
		logger  logging.Logger
		service services.ProductsService
	}
)

func NewProductsHandler(logger logging.Logger, service services.ProductsService) HTTPHandlers {
	return &productsHandler{logger, service}
}

func (h *productsHandler) Install(router server.HTTPServer) {
	router.Group("/v1/products", []*server.Route{
		server.NewRouteBuilder().Path("/").Method(http.MethodPost).Handler(h.post).Build(),
		server.NewRouteBuilder().Path("/").Method(http.MethodGet).Handler(h.list).Build(),
	})
}

// CreateProduct
// @Summary      CreateProducts
// @Description  Create a new Product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param				 product	body	views.CreateProductRequest	true	"Add Product"
// @Success      201  {object}  views.Product
// @Failure      400  {object}  views.HTTPError
// @Failure      404  {object}  views.HTTPError
// @Failure      500  {object}  views.HTTPError
// @Router       /products [post]
func (h *productsHandler) post(w http.ResponseWriter, req *http.Request) {
	body, err := ExtractBody[views.CreateProductRequest](w, req)
	if err != nil {
		h.logger.Error("unformatted body", zap.Error(err))
		return
	}

	grpcResponse, err := h.service.Create(req.Context(), &protos.CreateProductRequest{
		Name:     body.Name,
		Category: protos.Category(body.Category),
		Price:    body.Price,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(views.ProductFromProto(grpcResponse.Product).ToBuffer())
	w.WriteHeader(http.StatusCreated)
}

// ListProducts
// @Summary      ListProducts
// @Description  List Products with pagination
// @Tags         products
// @Accept       json
// @Produce      json
// @Param				 category	query	    string             true "Product Category" Enum(A, B, C)
// @Param        limit    query     int                true "Query Limit"      default(10)
// @Param        offset   query     int                true "Query Offset"     default(0)
// @Success      201      {object}  views.Product
// @Failure      400      {object}  views.HTTPError
// @Failure      404      {object}  views.HTTPError
// @Failure      500      {object}  views.HTTPError
// @Router       /products [get]
func (h *productsHandler) list(w http.ResponseWriter, req *http.Request) {
	body, err := ExtractBody[views.ListProductsRequest](w, req)
	if err != nil {
		h.logger.Error("unformatted body", zap.Error(err))
		return
	}

	grpcResponse, err := h.service.List(req.Context(), &protos.ListProductsRequest{
		Limit:      body.Limit,
		Offset:     body.Offset,
		Category:   protos.Category(body.Category),
		Ordination: (*protos.Ordination)(&body.Ordination),
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	list := []*views.Product{}
	for _, v := range grpcResponse.Products {
		list = append(list, views.ProductFromProto(v))
	}

	bytes, _ := json.Marshal(list)
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
