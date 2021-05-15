package services

import (
	"context"
	"github.com/oscarneiva/apilesson/domain/entities"
)

type UserService interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	GetByID(ctx context.Context, id string) (*entities.User, error)
	Search(ctx context.Context, query ...string) ([]*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
	Delete(ctx context.Context, id string) (*entities.User, error)
}