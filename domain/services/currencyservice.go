package services

import (
	"context"
	"github.com/oscarneiva/apilesson/domain/entities"
)

type CurrencyService interface {
	Create(ctx context.Context, user *entities.Currency) (*entities.Currency, error)
	GetByID(ctx context.Context, id string) (*entities.Currency, error)
	Search(ctx context.Context, query ...string) ([]*entities.Currency, error)
	Update(ctx context.Context, user *entities.Currency) (*entities.Currency, error)
	Delete(ctx context.Context, id string) (*entities.Currency, error)
}