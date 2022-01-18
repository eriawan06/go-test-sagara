package service

import (
	"context"

	"github.com/eriawan06/go-test-sagara/src/modules/product/model/domain"
	"github.com/eriawan06/go-test-sagara/src/modules/product/model/dto"
	"github.com/eriawan06/go-test-sagara/src/modules/product/repository"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	Repository repository.ProductRepository
	Validate   *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		Repository: productRepository,
		Validate:   validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request dto.ProductCreateDto) (dto.ProductResponseDto, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return dto.ProductResponseDto{}, err
	}

	product := domain.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Image:       request.Image,
	}

	product, err = service.Repository.Save(ctx, product)
	if err != nil {
		return dto.ProductResponseDto{}, err
	}

	return ToResponseDto(product), nil
}

func (service *ProductServiceImpl) Update(ctx context.Context, request dto.ProductUpdateDto) (dto.ProductResponseDto, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return dto.ProductResponseDto{}, err
	}

	product, err := service.Repository.FindById(ctx, request.Id)
	if err != nil {
		return dto.ProductResponseDto{}, err
	}

	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.Image = request.Image

	product, err = service.Repository.Update(ctx, product)
	if err != nil {
		return dto.ProductResponseDto{}, err
	}

	return ToResponseDto(product), nil
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId int) error {
	product, err := service.Repository.FindById(ctx, productId)
	if err != nil {
		return err
	}

	err = service.Repository.Delete(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) (dto.ProductResponseDto, error) {
	product, err := service.Repository.FindById(ctx, productId)
	if err != nil {
		return dto.ProductResponseDto{}, err
	}

	return ToResponseDto(product), nil
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) ([]dto.ProductResponseDto, error) {
	products, err := service.Repository.FindAll(ctx)
	// helper.PanicIfError(err)
	if err != nil {
		return nil, err
	}

	return ToResponsesDto(products), nil
}

func ToResponseDto(product domain.Product) dto.ProductResponseDto {
	return dto.ProductResponseDto{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ToResponsesDto(products []domain.Product) []dto.ProductResponseDto {
	var productResponsesDto []dto.ProductResponseDto
	for _, product := range products {
		productResponsesDto = append(productResponsesDto, ToResponseDto(product))
	}

	return productResponsesDto
}
