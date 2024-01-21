package handlers

import (
	"net/http"

	"github.com/ralvescosta/gokit/httpw/server"
	"github.com/ralvescosta/gokit/logging"
)

type (
	HTTPHandlers interface {
		Install(router server.HTTPServer)
	}

	productsHandler struct {
		logger logging.Logger
	}
)

func NewProductsHandler(logger logging.Logger) HTTPHandlers {
	return &productsHandler{logger}
}

func (h *productsHandler) Install(router server.HTTPServer) {
	router.Group("/products", []*server.Route{
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
// @Param				 product	body		httpHandlers.Product	true	"Add Product"
// @Success      201  {object}  httpHandlers.Product
// @Failure      400  {object}  httpHandlers.HTTPError
// @Failure      404  {object}  httpHandlers.HTTPError
// @Failure      500  {object}  httpHandlers.HTTPError
// @Router       /products/ [post]
func (h *productsHandler) post(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("post")

	w.WriteHeader(200)
}

// ListProducts
// @Summary      ListProducts
// @Description  List Products with pagination
// @Tags         products
// @Accept       json
// @Produce      json
// @Param				 category	path	    int                     true	"Add Product"
// @Success      201      {object}  httpHandlers.Product
// @Failure      400      {object}  httpHandlers.HTTPError
// @Failure      404      {object}  httpHandlers.HTTPError
// @Failure      500      {object}  httpHandlers.HTTPError
// @Router       /products/{id} [post]
func (h *productsHandler) list(w http.ResponseWriter, req *http.Request) {
	h.logger.Info("post")

	w.WriteHeader(200)
}
