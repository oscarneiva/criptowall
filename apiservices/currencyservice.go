package apiservices

import (
	"context"
	"errors"
	"github.com/oscarneiva/apilesson/domain/entities"
	"github.com/oscarneiva/apilesson/domain/repositories"
)

type currencyService struct {
	repository repositories.CurrencyRepository
}

func NewCurrencyService(repository repositories.CurrencyRepository) *currencyService {
	return &currencyService{repository: repository}
}

func (self *currencyService) Create(ctx context.Context, currency *entities.Currency) (*entities.Currency, error) {
	existCurrency, err := self.repository.GetByID(ctx, currency.ID)
	if err != nil{
		return nil, err
	}

	if existCurrency != nil{
		return nil, errors.New("Error: Currency already exist!")
	}

	err = self.repository.Create(ctx, currency)
	if err != nil{
		return nil, err
	}

	return currency, nil
}

func (self *currencyService) GetByID(ctx context.Context, id string) (*entities.Currency, error) {
	currency, err := self.repository.GetByID(ctx, id)
	if err != nil{
		return nil, err
	}
	return currency, nil
}

func (self *currencyService) Search(ctx context.Context, query ...string) ([]*entities.Currency, error) {
	currencies, err := self.repository.Search(ctx, query...)
	if err != nil{
		return nil, err
	}
	return currencies, nil
}

func (self *currencyService) Update(ctx context.Context, currency *entities.Currency) (*entities.Currency, error) {
	_, err := self.repository.GetByID(ctx, currency.ID)
	if err != nil{
		return nil, err
	}

	err = self.repository.Update(ctx, currency)
	if err != nil {
		return nil, err
	}

	return currency, nil
}

func (self *currencyService) Delete(ctx context.Context, id string) (*entities.Currency, error) {
	self.repository.Delete(ctx, id)
	return nil, nil
}

