package usecase

import (
	"auth-golang/internal/entity"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repository UserRepo
}

func New(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repository: r,
	}
}

func (c *UserUseCase) Create(ctx context.Context, user entity.User) error {
	if c.repository.IsExistsEmail(ctx, user.Email) {
		return fmt.Errorf("input data error: email already exists")
	}
	password, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(password)
	err = c.repository.AddUser(ctx, user)
	if err != nil {
		return err
	}
	return nil

}

func (c *UserUseCase) Get(ctx context.Context, id string) (*entity.User, error) {
	user, err := c.repository.GetUser(ctx, id)
	return user, err
}

func hashPassword(password string) ([]byte, error) {
	pass := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}
