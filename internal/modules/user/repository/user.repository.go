package repository

import (
	"context"

	"vibe-user/internal/modules/user/entity"
	"vibe-user/pkg/abstract"

	"gorm.io/gorm"
)

type UserRepository interface {
	abstract.IAbstractRepository[entity.User]
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type userRepository struct {
	*abstract.AbstractRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{abstract.NewAbstractRepository[entity.User](db)}
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.GetDB().WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
