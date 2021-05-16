package apirepositories

import (
	"context"
	"github.com/oscarneiva/apilesson/domain/entities"
)

type walletCSVRepository struct {
	wallets map[string] *entities.Wallet
}

func NewWalletCSVRepository() *walletCSVRepository  {
	return &walletCSVRepository{wallets: make(map[string] *entities.Wallet)}
}

func (self *walletCSVRepository) Create(ctx context.Context, wallet *entities.Wallet) error {
	self.wallets[wallet.ID] = wallet
	return nil
}

func (self *walletCSVRepository) GetByID(ctx context.Context, id string) (*entities.Wallet, error) {
	return self.wallets[id], nil
}

func (self *walletCSVRepository) Search(ctx context.Context, query ...string) ([]*entities.Wallet, error) {
	var wallets []*entities.Wallet
	for _, wallet := range self.wallets{
		wallets = append(wallets, wallet)
	}
	return wallets,nil
}

func (self *walletCSVRepository) Update(ctx context.Context, wallet *entities.Wallet) error {
	self.wallets[wallet.ID] = wallet
	return nil
}

func (self *walletCSVRepository) Delete(ctx context.Context, id string) error {
	self.wallets[id] = nil
	return nil
}
