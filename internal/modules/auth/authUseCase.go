package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gutkedu/golang_api/internal/modules/user"
	"golang.org/x/crypto/bcrypt"
)

// Implementation of the repository in this service.
type authUserUseCase struct {
	userRepository user.UserRepository
}

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewAuthUserUseCase(r user.UserRepository) AuthUserUseCase {
	return &authUserUseCase{
		userRepository: r,
	}
}

func (s *authUserUseCase) Execute(
	ctx context.Context,
	auth AuthRequest) (AuthResponse, error) {

	//get user
	user, err := s.userRepository.FindByEmail(ctx, auth.Email)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("email or password is incorrect!")
	}
	//check password hash
	if !CheckPasswordHash(auth.Password, user.Password) {
		return AuthResponse{}, fmt.Errorf("email or password is incorrect")
	}
	//generate new token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("Secret"))
	if err != nil {
		return AuthResponse{}, err
	}
	return AuthResponse{*user, t}, err
}
