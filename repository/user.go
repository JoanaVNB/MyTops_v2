package repository

import (
	"context"
	"app/domain"
)

type UserRepository interface {
	Create(context.Context, domain.User) (domain.User, error)
	GetID(context.Context, string, domain.User) (domain.User, error)
	Login(context.Context, domain.User, domain.Login) bool
} 

