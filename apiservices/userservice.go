package apiservices

import (
	"context"
	"errors"
	"github.com/oscarneiva/apilesson/domain/entities"
	"github.com/oscarneiva/apilesson/domain/repositories"
)

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository: repository}
}

func (self *userService) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	existUser, err := self.repository.GetByID(ctx, user.ID)
	if err != nil{
		return nil, err
	}

	if existUser != nil{
		return nil, errors.New("Error: User already exist!")
	}

	err = self.repository.Create(ctx, user)
	if err != nil{
		return nil, err
	}

	return user, nil
}

func (self *userService) GetByID(ctx context.Context, id string) (*entities.User, error) {
	user, err := self.repository.GetByID(ctx, id)
	if err != nil{
		return nil, err
	}
	return user, nil
}

func (self *userService) Search(ctx context.Context, query ...string) ([]*entities.User, error) {
	users, err := self.repository.Search(ctx, query...)
	if err != nil{
		return nil, err
	}
	return users, nil
}

func (self *userService) Update(ctx context.Context, user *entities.User) (*entities.User, error) {
	_, err := self.repository.GetByID(ctx, user.ID)
	if err != nil{
		return nil, err
	}

	err = self.repository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (self *userService) Delete(ctx context.Context, id string) (*entities.User, error) {
	self.repository.Delete(ctx, id)
	return nil, nil
}

