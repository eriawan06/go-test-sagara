package repository

import (
	"context"
	"errors"

	"github.com/eriawan06/go-test-sagara/src/modules/user/model/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryImpl struct {
	DB *sqlx.DB
}

func NewUserUserRepository(DB *sqlx.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (repo *UserRepositoryImpl) Save(ctx context.Context, user domain.User) (domain.User, error) {
	query := `
	INSERT INTO user(fullname, username, password)
	VALUES (?, ?, ?)
	`
	result, err := repo.DB.ExecContext(ctx, query, user.Fullname, user.Username, user.Password)
	if err != nil {
		return domain.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.User{}, err
	}

	insertedUser, _ := repo.FindById(ctx, int(id))
	user = insertedUser
	return user, nil
}

// Update(ctx context.Context, user domain.User) (domain.User, error)
// Delete(ctx context.Context, user domain.User) error

func (repo *UserRepositoryImpl) FindById(ctx context.Context, userId int) (domain.User, error) {
	query := `SELECT * FROM user WHERE id = ?`
	row := repo.DB.QueryRowxContext(ctx, query, userId)

	user := domain.User{}

	if row == nil {
		return user, errors.New("user is not found")
	}

	err := row.StructScan(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	query := `SELECT * FROM user WHERE username = ?`
	row := repo.DB.QueryRowxContext(ctx, query, username)

	user := domain.User{}

	if row == nil {
		return user, errors.New("user is not found")
	}

	err := row.StructScan(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repo *UserRepositoryImpl) FindAll(ctx context.Context) ([]domain.User, error) {
	query := `SELECT * FROM user`
	rows, err := repo.DB.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err = rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
