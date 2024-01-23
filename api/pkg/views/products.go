package views

import (
	"encoding/json"
	"time"

	"github.com/ralvescosta/ec2-hello-world/protos"
	"github.com/ralvescosta/gokit/guid"
)

type (
	Category   int8
	Ordination int32

	// Product
	Product struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Category  Category  `json:"category"`
		Price     float32   `json:"price"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
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

func (r *CreateProductRequest) ToProto() *protos.CreateProductRequest {
	return &protos.CreateProductRequest{
		Name:     r.Name,
		Category: protos.Category(r.Category),
		Price:    r.Price,
	}
}

func (r *ListProductsRequest) ToProto() *protos.ListProductsRequest {
	return &protos.ListProductsRequest{
		Limit:      r.Limit,
		Offset:     r.Offset,
		Category:   protos.Category(r.Category),
		Ordination: (*protos.Ordination)(&r.Ordination),
	}
}

func ProductFromProto(p *protos.Product) *Product {
	return &Product{
		ID:        guid.StringFromByteSliceUUID(p.Id),
		Name:      p.Name,
		Category:  Category(p.Category),
		Price:     p.Price,
		CreatedAt: p.CreatedAt.AsTime(),
		UpdatedAt: p.UpdatedAt.AsTime(),
	}
}

func (p *Product) ToBuffer() []byte {
	bytes, _ := json.Marshal(p)
	return bytes
}
