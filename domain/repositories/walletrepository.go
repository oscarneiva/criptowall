package repositories

import (
	"context"
	"github.com/oscarneiva/apilesson/domain/entities"
)

type WalletRepository interface {
	Create(ctx context.Context, user *entities.Wallet) error
	GetByID(ctx context.Context, id string) (*entities.Wallet, error)
	Search(ctx context.Context, query ...string) ([]*entities.Wallet, error)
	Update(ctx context.Context, user *entities.Wallet) error
	Delete(ctx context.Context, id string) error
}