package service

import (
	"context"
	"errors"

	"github.com/eriawan06/go-test-sagara/src/helper"
	"github.com/eriawan06/go-test-sagara/src/modules/auth/model/dto"
	userDto "github.com/eriawan06/go-test-sagara/src/modules/user/model/dto"
	us "github.com/eriawan06/go-test-sagara/src/modules/user/service"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	UserService us.UserService
	Validate    *validator.Validate
}

func NewAuthService(userService us.UserService, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserService: userService,
		Validate:    validate,
	}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request dto.AuthLoginDto) (dto.AuthLoginResponseDto, error) {
	user, err := service.UserService.FindByUsername(ctx, request.Username)
	if err != nil {
		return dto.AuthLoginResponseDto{}, err
	}

	isPasswordValid := helper.CheckPasswordHash(request.Password, user.Password)
	if !isPasswordValid {
		err = errors.New("wrong username or password")
		return dto.AuthLoginResponseDto{}, err
	}

	tokens, _, _ := helper.GenerateJWT(user.Username)
	loginResponse := dto.AuthLoginResponseDto{
		Tokens: tokens,
		User:   us.ToResponseDtoRemovePassword(user),
	}

	return loginResponse, nil
}

func (service *AuthServiceImpl) Register(ctx context.Context, request userDto.UserCreateDto) (userDto.UserResponseDto, error) {
	user, err := service.UserService.Create(ctx, request)
	if err != nil {
		return user, err
	}

	return user, nil
}
