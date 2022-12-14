package service

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import (
	"github.com/google/uuid"
	"context"
	"app/domain"
	"app/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService (repository repository.UserRepository) *UserService{
	return &UserService{repository: repository}
}

func (us UserService) Create(ctx context.Context, u domain.User) (user domain.User, err error){
	u.ID = uuid.NewString()
	us.repository.Create(ctx,u)
	return u, err
}

func (us UserService) GetID(ctx context.Context, id string, u domain.User)(user domain.User, err error){
	user, err = us.repository.GetID(ctx, id, u)
	if err != nil{
		return user, err
	}
	return user, nil
}