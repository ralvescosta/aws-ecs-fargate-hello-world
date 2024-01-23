package handlers

import (
	"net/http"

	"github.com/ralvescosta/ecs-hello-world/api/pkg/internal/services"
	"github.com/ralvescosta/ecs-hello-world/api/pkg/views"
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

	println(body)

	w.WriteHeader(200)
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
	h.logger.Info("get")

	w.WriteHeader(200)
}
