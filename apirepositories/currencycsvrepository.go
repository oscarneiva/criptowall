package apirepositories

import (
	"context"
	"github.com/oscarneiva/apilesson/domain/entities"
)

type currencyCSVRepository struct {
	currencies map[string] *entities.Currency
}

func NewCurrencyCSVRepository() *currencyCSVRepository {
	return &currencyCSVRepository{currencies: make(map[string] *entities.Currency)}
}

func (self *currencyCSVRepository) Create(ctx context.Context, currency *entities.Currency) error {
	self.currencies[currency.ID] = currency
	return nil
}

func (self *currencyCSVRepository) GetByID(ctx context.Context, id string) (*entities.Currency, error) {
	return self.currencies[id], nil
}

func (self *currencyCSVRepository) Search(ctx context.Context, query ...string) ([]*entities.Currency, error) {
	var currencies []*entities.Currency
	for _, currency := range self.currencies {
		currencies = append(currencies, currency)
	}
	return currencies, nil
}

func (self *currencyCSVRepository) Update(ctx context.Context, currency *entities.Currency) error {
	self.currencies[currency.ID] = currency
	return nil
}

func (self *currencyCSVRepository) Delete(ctx context.Context, id string) error {
	self.currencies[id] = nil
	return nil
}
