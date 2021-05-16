package apiservices

import (
	"context"
	"errors"
	"github.com/oscarneiva/apilesson/domain/entities"
	"github.com/oscarneiva/apilesson/domain/repositories"
)

type walletService struct {
	repository repositories.WalletRepository
}

func NewWalletService(repository repositories.WalletRepository) *walletService {
	return &walletService{repository: repository}
}

func (self *walletService) Create(ctx context.Context, wallet *entities.Wallet) (*entities.Wallet, error) {
	existWallet, err := self.repository.GetByID(ctx, wallet.ID)
	if err != nil{
		return nil, err
	}

	if existWallet != nil{
		return nil, errors.New("Error: User already exist!")
	}

	err = self.repository.Create(ctx, wallet)
	if err != nil{
		return nil, err
	}

	return wallet, nil
}

func (self *walletService) GetByID(ctx context.Context, id string) (*entities.Wallet, error) {
	wallet, err := self.repository.GetByID(ctx, id)
	if err != nil{
		return nil, err
	}
	return wallet, nil
}

func (self *walletService) Search(ctx context.Context, query ...string) ([]*entities.Wallet, error) {
	wallets, err := self.repository.Search(ctx, query...)
	if err != nil{
		return nil, err
	}
	return wallets, nil
}

func (self *walletService) Update(ctx context.Context, wallet *entities.Wallet) (*entities.Wallet, error) {
	_, err := self.repository.GetByID(ctx, wallet.ID)
	if err != nil{
		return nil, err
	}

	err = self.repository.Update(ctx, wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (self *walletService) Delete(ctx context.Context, id string) (*entities.Wallet, error) {
	self.repository.Delete(ctx, id)
	return nil, nil
}
