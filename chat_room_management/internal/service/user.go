package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"ahyalfan.my.id/chat_rom_management/domain"
	"ahyalfan.my.id/chat_rom_management/dto"
	"ahyalfan.my.id/chat_rom_management/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	cnf            *config.Config
	userRepository domain.UserRepository
}

func NewUser(cnf *config.Config, userRepository domain.UserRepository) domain.UserService {
	return &userService{userRepository: userRepository, cnf: cnf}
}

// CreateUser implements domain.UserService.
func (u *userService) CreateUser(ctx context.Context, req dto.UserCreatedReq) (dto.UserCreatedRes, error) {
	result, _ := u.userRepository.FindByEmail(ctx, req.Email)
	if result.ID > 0 {
		fmt.Println(result)
		return dto.UserCreatedRes{}, domain.ErrEmailTaken
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	err := u.userRepository.Create(ctx, user)
	if err != nil {
		return dto.UserCreatedRes{}, err
	}
	return dto.UserCreatedRes{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

// LoginUser implements domain.UserService.
func (u *userService) LoginUser(ctx context.Context, req dto.LoginUserReq) (dto.LoginUserRes, error) {
	user, err := u.userRepository.FindByEmail(ctx, req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.LoginUserRes{}, domain.ErrAuthFailed
	}
	if err != nil {
		return dto.LoginUserRes{}, domain.ErrAuthFailed
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.LoginUserRes{}, domain.ErrAuthFailed
	}

	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Minute * time.Duration(u.cnf.JWT.Expired)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(u.cnf.JWT.Key))

	if err != nil {
		return dto.LoginUserRes{}, err
	}
	return dto.LoginUserRes{
		ID:          user.ID,
		Username:    user.Username,
		AccessToken: tokenString,
	}, nil
}
