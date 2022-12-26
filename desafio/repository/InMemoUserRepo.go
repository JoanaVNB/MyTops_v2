package repository

import (
	"app/domain"
	"context"
	"fmt"
)

type InMemoUserRepository struct{
	uMap	map[string]domain.User
}

func NewInMemoUserRepository() *InMemoUserRepository{
	return &InMemoUserRepository{uMap: make(map[string]domain.User)}
}

func (in InMemoUserRepository) Create(c context.Context, u domain.User) (domain.User, error){
	in.uMap[u.ID] = u
	fmt.Println("Criado")
	return u,  nil
}

func (in InMemoUserRepository) GetID(c context.Context, id string, u domain.User) (domain.User, error){
	for _, k := range in.uMap{
		gotID := k.ID
		if gotID == id{
		return u, nil
		}
	}
	return u, nil
}

func (in InMemoUserRepository) Login(c context.Context, u domain.User, l domain.Login) bool{
	email := l.Email
	
	for _, k := range in.uMap{
		gotEmail := k.Email
		if gotEmail == email{
			u = k
		}
	}
	if l.Password == u.Password{
		return true
	}
	return false
}