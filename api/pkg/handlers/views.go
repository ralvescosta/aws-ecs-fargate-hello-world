package handlers

import "time"

type (
	Category   int8
	Ordination int32

	// HTTPError
	HTTPError struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	}

	// Product
	Product struct {
		ID        string      `json:"id"`
		Name      string      `json:"name"`
		Category  interface{} `json:"category"`
		Price     float32     `json:"price"`
		CreatedAt time.Time   `json:"createdAt"`
		UpdatedAt time.Time   `json:"updatedAt"`
	}

	// CreateProductRequest
	CreateProductRequest struct {
		Name     string   `json:"name"     validate:"required,min=3"   example:"name" binding:"required"`
		Category Category `json:"category" validate:"required"         example:"1"    binding:"required"`
		Price    float32  `json:"price"    validate:"required,gte=1.0" example:"1.2"  binding:"required"`
	}

	// ListProductsRequest
	ListProductsRequest struct {
		Limit      *int32     `json:"limit"       validate:"min=1"`
		Offset     *int32     `json:"offset"      validate:"min=1"`
		Category   Category   `json:"category"    validate:"required"  binding:"required"`
		Ordination Ordination `json:"ordination"`
	}
)

var (
	UnknownCategory Category = 0
	CategoryA       Category = 1
	CategoryB       Category = 2
	CategoryC       Category = 3

	UnknownOrdination Ordination = 0
	OrdinationAsc                = 1
	OrdinationDesc               = 2
)
