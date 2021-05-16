package services

import (
	"context"
	"github.com/oscarneiva/apilesson/domain/entities"
)

type WalletService interface {
	Create(ctx context.Context, wallet *entities.Wallet) (*entities.Wallet, error)
	GetByID(ctx context.Context, id string) (*entities.Wallet, error)
	Search(ctx context.Context, query ...string) ([]*entities.Wallet, error)
	Update(ctx context.Context, user *entities.Wallet) (*entities.Wallet, error)
	Delete(ctx context.Context, id string) (*entities.Wallet, error)
}