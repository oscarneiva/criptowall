package apirepositories

import (
	"context"
	"github.com/oscarneiva/apilesson/domain/entities"
)

type userCSVRepository struct {

}

func (self *userCSVRepository) Create(ctx context.Context, user *entities.User) error {
	
	return nil
}

func (self *userCSVRepository) GetByID(ctx context.Context, id string) (*entities.User, error) {
	return nil, nil
}

func (self *userCSVRepository) Search(ctx context.Context, query ...string) ([]*entities.User, error) {
	return nil, nil
}

func (self *userCSVRepository) Update(ctx context.Context, user *entities.User) error {
	return nil
}

func (self *userCSVRepository) Delete(ctx context.Context, id string) error {
	return nil
}
