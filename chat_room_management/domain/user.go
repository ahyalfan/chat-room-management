package domain

import (
	"context"

	"ahyalfan.my.id/chat_rom_management/dto"
)

type User struct {
	ID       int64  `gorm:"primary_key;autoIncrement"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, req dto.UserCreatedReq) (dto.UserCreatedRes, error)
	LoginUser(ctx context.Context, req dto.LoginUserReq) (dto.LoginUserRes, error)
}
