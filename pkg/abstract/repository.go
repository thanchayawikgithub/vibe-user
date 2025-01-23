package abstract

import (
	"context"
	"log"

	"vibe-user/internal/modules/user/entity"

	"gorm.io/gorm"
)

type IEntity interface {
	entity.User
}

type IAbstractRepository[T IEntity] interface {
	FindByID(ctx context.Context, id string) (*T, error)
	Create(ctx context.Context, entity *T) (*T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, entity *T) error
}

type AbstractRepository[T IEntity] struct {
	db *gorm.DB
}

func NewAbstractRepository[T IEntity](db *gorm.DB) *AbstractRepository[T] {
	return &AbstractRepository[T]{db}
}

func (r *AbstractRepository[T]) GetDB() *gorm.DB {
	return r.db
}

func (r *AbstractRepository[T]) Create(ctx context.Context, entity *T) (*T, error) {
	if err := r.db.WithContext(ctx).Create(entity).Error; err != nil {
		log.Println("Error creating entity: %v", err)
		return nil, err
	}
	return entity, nil
}

func (r *AbstractRepository[T]) FindByID(ctx context.Context, id string) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).First(&entity, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *AbstractRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *AbstractRepository[T]) Delete(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Delete(entity).Error
}
