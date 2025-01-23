package service

import (
	"context"
	"vibe-user/internal/config"
	"vibe-user/internal/modules/user/entity"
	"vibe-user/internal/modules/user/repository"
	"vibe-user/pkg/jwt"

	"gorm.io/gorm"
)

type AuthService interface {
	Login(ctx context.Context, user *entity.User) (string, error)
}

type authService struct {
	userRepository repository.UserRepository
	cfg            *config.Jwt
}

func NewAuthService(userRepository repository.UserRepository, cfg *config.Jwt) AuthService {
	return &authService{userRepository, cfg}
}

func (s *authService) Login(ctx context.Context, user *entity.User) (string, error) {
	var loginUser *entity.User
	loginUser, err := s.userRepository.FindByEmail(ctx, user.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			loginUser, err := s.userRepository.Create(ctx, user)
			if err != nil {
				return "", err
			}

			token, err := jwt.GenerateToken(s.cfg, loginUser)
			if err != nil {
				return "", err
			}

			return token, nil
		}
		return "", err
	}

	token, err := jwt.GenerateToken(s.cfg, loginUser)
	if err != nil {
		return "", err
	}

	return token, nil
}
