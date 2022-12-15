package presenters

import (
	"app/domain"
)

type User struct {
	Name     string
	Email    string
}

type Login struct{
	Email    string
	Password string
}

func PresenterUser(u domain.User) *User{
	return &User{
		Name: u.Name,
		Email: u.Email,
	}
}
