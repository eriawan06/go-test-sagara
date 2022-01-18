package dto

import (
	"github.com/eriawan06/go-test-sagara/src/modules/auth/model/domain"
	user "github.com/eriawan06/go-test-sagara/src/modules/user/model/dto"
)

type AuthLoginResponseDto struct {
	Tokens domain.Jwt           `json:"token"`
	User   user.UserResponseDto `json:"user"`
}
