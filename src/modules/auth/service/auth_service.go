package service

import (
	"context"

	"github.com/eriawan06/go-test-sagara/src/modules/auth/model/dto"
	userDto "github.com/eriawan06/go-test-sagara/src/modules/user/model/dto"
)

type AuthService interface {
	Login(ctx context.Context, request dto.AuthLoginDto) (dto.AuthLoginResponseDto, error)
	Register(ctx context.Context, request userDto.UserCreateDto) (userDto.UserResponseDto, error)
}
