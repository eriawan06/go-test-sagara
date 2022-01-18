package repository

import (
	"context"

	"github.com/eriawan06/go-test-sagara/src/modules/user/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
	// Update(ctx context.Context, user domain.User) (domain.User, error)
	// Delete(ctx context.Context, user domain.User) error
	FindById(ctx context.Context, userId int) (domain.User, error)
	FindByUsername(ctx context.Context, username string) (domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
}
