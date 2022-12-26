package service

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import (
	"app/domain"
	"context"
	"github.com/google/uuid"
)
type UserRepository interface {
	Create(context.Context, domain.User) (domain.User, error)
	GetID(context.Context, string, domain.User) (domain.User, error)
	Login(context.Context, domain.User, domain.Login) bool
} 

type UserService struct {
	repository UserRepository
}

func NewUserService (repository UserRepository) *UserService{
	return &UserService{repository: repository}
}

//também é validado por ShouldBindJSON no package controller
func validate(u domain.User) bool{
	if u.Name=="" || u.Email=="" || u.Password==""{
		return false
	}
	return true
}

//Create retorna user apenas para verificar se foi criado ou não, pois se foi criado irá retornar com valores, caso não tenha sido criado, irá retornar vazio
func (us UserService) Create(ctx context.Context, u domain.User) (domain.User, error){
	if validate(u) == true {
		u.ID = uuid.NewString()
		user, err := us.repository.Create(ctx,u)
		user.ID = u.ID
		if err != nil{
			return user, err
		}
		return user, nil
	}
	return domain.User{}, nil
}

func (us UserService) GetID(ctx context.Context, id string, u domain.User)(domain.User, error){
	user, err := us.repository.GetID(ctx, id, u)
	if err != nil{
		return domain.User{}, err
	}
	return user, nil
}

func (us UserService) Login(c context.Context, u domain.User, l domain.Login) bool{
	return us.repository.Login(c, u, l)
}