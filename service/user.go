package service

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import (
	"github.com/google/uuid"
	"context"
	"app/domain"
)

type UserServiceRepository interface {
	Create (context.Context, domain.User) error
	GetID (context.Context, string, domain.User) (domain.User, error)
}

type UserService struct {
	repository UserServiceRepository
}

func NewUserUseCase (repository UserServiceRepository) *UserService{
	return &UserService{repository: repository}
}

func (user UserService) Create (ctx context.Context, u domain.User) (domain.User, error){
	u.ID = uuid.NewString()

	err := user.repository.Create(ctx, u)
	if err != nil{
		return domain.User{}, err
	}
	return u, nil
}

func (user UserService) GetID (ctx context.Context, id string, u domain.User)(domain.User, error){

	u, err := user.repository.GetID(ctx, id, u)
	if err != nil{
		return domain.User{}, err
	}
	return u, nil
}