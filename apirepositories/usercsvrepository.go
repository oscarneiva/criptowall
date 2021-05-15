package apirepositories

import (
	"context"
	"github.com/oscarneiva/apilesson/domain/entities"
)

type userCSVRepository struct {
	users map[string] *entities.User
}

func NewUserCSVRepository() *userCSVRepository {
	return &userCSVRepository{users: make(map[string] *entities.User)}
}


func (self *userCSVRepository) Create(ctx context.Context, user *entities.User) error {
	self.users[user.ID] = user
	return nil
}

func (self *userCSVRepository) GetByID(ctx context.Context, id string) (*entities.User, error) {
	return self.users[id], nil
}

func (self *userCSVRepository) Search(ctx context.Context, query ...string) ([]*entities.User, error) {
	var users []*entities.User
	for _, user := range self.users {
		users = append(users, user)
	}
	return users, nil
}

func (self *userCSVRepository) Update(ctx context.Context, user *entities.User) error {
	self.users[user.ID] = user
	return nil
}

func (self *userCSVRepository) Delete(ctx context.Context, id string) error {
	self.users[id] = nil
	return nil
}
