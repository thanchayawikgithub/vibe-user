package service

import (
	"context"
	"vibe-user/internal/modules/user/entity"
	"vibe-user/internal/modules/user/repository"

	"gorm.io/gorm"
)

type AuthService interface {
	Login(ctx context.Context, user *entity.User) (*entity.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository}
}

func (s *authService) Login(ctx context.Context, user *entity.User) (*entity.User, error) {
	existingUser, err := s.userRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return s.userRepository.Create(ctx, user)
		}
		return nil, err
	}

	return existingUser, nil
}
