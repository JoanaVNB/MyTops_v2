package repository

import (
	"context"
	"app/domain"
)

type InMemoUserRepository struct{
	u	map[string]domain.User
}

func NewInMemoUserRepository() *InMemoUserRepository{
	return &InMemoUserRepository{u: make(map[string]domain.User)}
}

func (in InMemoUserRepository) Create(ctx context.Context, u domain.User) error{
	in.u[u.ID] = u
	return nil
}

func (in InMemoUserRepository) GetID(ctx context.Context, id string, u domain.User) (domain.User, error){
	for k, _ := range in.u{
		gotID := k
		if gotID == id{
		return u, nil
		}
	}
	return u, nil
}

