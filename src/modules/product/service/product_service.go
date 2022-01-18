package service

import (
	"context"

	"github.com/eriawan06/go-test-sagara/src/modules/product/model/dto"
)

type ProductService interface {
	Create(ctx context.Context, request dto.ProductCreateDto) (dto.ProductResponseDto, error)
	Update(ctx context.Context, request dto.ProductUpdateDto) (dto.ProductResponseDto, error)
	Delete(ctx context.Context, productId int) error
	FindById(ctx context.Context, productId int) (dto.ProductResponseDto, error)
	FindAll(ctx context.Context) ([]dto.ProductResponseDto, error)
}
