package presenters

import (
	"app/domain"
)

type User struct {
	ID	string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Login struct{
	Email    string `json:"email"`
	Password string `json:"password"`
}

func PresenterUser(u domain.User) *User{
	return &User{
		ID: u.ID,
		Name: u.Name,
		Email: u.Email,
	}
}