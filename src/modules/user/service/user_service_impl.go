package service

import (
	"context"

	"github.com/eriawan06/go-test-sagara/src/helper"
	"github.com/eriawan06/go-test-sagara/src/modules/user/model/domain"
	"github.com/eriawan06/go-test-sagara/src/modules/user/model/dto"
	"github.com/eriawan06/go-test-sagara/src/modules/user/repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	Repository repository.UserRepository
	Validate   *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		Repository: userRepository,
		Validate:   validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request dto.UserCreateDto) (dto.UserResponseDto, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return dto.UserResponseDto{}, err
	}

	hashedPassword, _ := helper.HashPassword(request.Password)
	user := domain.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Password: hashedPassword,
	}

	user, err = service.Repository.Save(ctx, user)
	if err != nil {
		return dto.UserResponseDto{}, err
	}

	return ToResponseDto(user), nil
}

// Update(ctx context.Context, request dto.UserUpdateDto) (dto.UserResponseDto, error)
// Delete(ctx context.Context, userId int) error

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) (dto.UserResponseDto, error) {
	user, err := service.Repository.FindById(ctx, userId)
	if err != nil {
		return dto.UserResponseDto{}, err
	}

	return ToResponseDto(user), nil
}

func (service *UserServiceImpl) FindByUsername(ctx context.Context, username string) (dto.UserResponseWithPasswordDto, error) {
	user, err := service.Repository.FindByUsername(ctx, username)
	if err != nil {
		return dto.UserResponseWithPasswordDto{}, err
	}

	return ToResponseDtoWithPassword(user), nil
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]dto.UserResponseDto, error) {
	users, err := service.Repository.FindAll(ctx)
	// helper.PanicIfError(err)
	if err != nil {
		return nil, err
	}

	return ToResponsesDto(users), nil
}

func ToResponseDto(user domain.User) dto.UserResponseDto {
	return dto.UserResponseDto{
		Id:        user.Id,
		Fullname:  user.Fullname,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToResponseDtoRemovePassword(user dto.UserResponseWithPasswordDto) dto.UserResponseDto {
	return dto.UserResponseDto{
		Id:        user.Id,
		Fullname:  user.Fullname,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToResponseDtoWithPassword(user domain.User) dto.UserResponseWithPasswordDto {
	return dto.UserResponseWithPasswordDto{
		Id:        user.Id,
		Fullname:  user.Fullname,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToResponsesDto(users []domain.User) []dto.UserResponseDto {
	var userResponsesDto []dto.UserResponseDto
	for _, user := range users {
		userResponsesDto = append(userResponsesDto, ToResponseDto(user))
	}

	return userResponsesDto
}
