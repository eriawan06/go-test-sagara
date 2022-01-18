package dto

import (
	user "github.com/eriawan06/go-test-sagara/src/modules/user/model/dto"
	"github.com/eriawan06/go-test-sagara/src/shared"
)

type AuthLoginResponseDto struct {
	Tokens shared.JwtModel      `json:"token"`
	User   user.UserResponseDto `json:"user"`
}
