package service

import (
	"context"
	"vibe-user/internal/modules/user/entity"
	"vibe-user/internal/modules/user/repository"
)

type UserService interface {
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByID(ctx context.Context, id string) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Create(ctx context.Context, user *entity.User) error {
	return s.userRepository.Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, user *entity.User) error {
	return s.userRepository.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, user *entity.User) error {
	return s.userRepository.Delete(ctx, user)
}

func (s *userService) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	return s.userRepository.FindByEmail(ctx, email)
}

func (s *userService) FindByID(ctx context.Context, id string) (*entity.User, error) {
	return s.userRepository.FindByID(ctx, id)
}
