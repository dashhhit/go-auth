package usecase

import (
	"auth-golang/config"
	"auth-golang/internal/entity"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginUseCase struct {
	repository UserRepo
}

func NewLoginUseCase(r UserRepo) *LoginUseCase {
	return &LoginUseCase{
		repository: r,
	}
}

func (r *LoginUseCase) GenerateJWT(input *entity.LoginInput) (string, error) {

	user, err := r.repository.GetUserByEmail(context.Background(), input.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", err
	}

	t, err := generateJWT(user)
	if err != nil {
		return "", err
	}
	return t, nil

}

func generateJWT(user *entity.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["first_name"] = user.FirstName
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.GetConfig("SECRET")))
	if err != nil {
		return "", err
	}
	return t, nil
}
