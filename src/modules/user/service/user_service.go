package service

import (
	"context"

	"github.com/eriawan06/go-test-sagara/src/modules/user/model/dto"
)

type UserService interface {
	Create(ctx context.Context, request dto.UserCreateDto) (dto.UserResponseDto, error)
	// Update(ctx context.Context, request dto.UserUpdateDto) (dto.UserResponseDto, error)
	// Delete(ctx context.Context, userId int) error
	FindById(ctx context.Context, userId int) (dto.UserResponseDto, error)
	FindByUsername(ctx context.Context, username string) (dto.UserResponseWithPasswordDto, error)
	FindAll(ctx context.Context) ([]dto.UserResponseDto, error)
}
