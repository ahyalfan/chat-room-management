package repository

import (
	"context"

	"ahyalfan.my.id/chat_rom_management/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// Create implements domain.UserRepository.
func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	return u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(user).Error
	})
}

// FindByEmail implements domain.UserRepository.
func (u *userRepository) FindByEmail(ctx context.Context, email string) (users *domain.User, err error) {
	err = u.db.WithContext(ctx).Where("email = ?", email).First(&users).Error
	return
}
