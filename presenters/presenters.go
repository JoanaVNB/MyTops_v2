package presenters

import (
	"app/domain"
)

type User struct {
	ID	string
	Name     string
	Email    string
}

type Login struct{
	Email    string
	Password string
}

func PresenterUser(u domain.User) *User{
	return &User{
		ID: u.ID,
		Name: u.Name,
		Email: u.Email,
	}
}

type Shop struct {
	ID      string  
	Name    string  
	Flavors [6]string 
	Score   float64 
	Price   float64 
	Link    string  
	Favorite	bool 
}

type NameUpdated struct{
	NewName	string	
}

type PriceUpdated struct{
	NewPrice	float64	
}

func PresenterShop(s domain.Shop) *Shop{
	return &Shop{
		ID: s.ID,
		Name: s.Name,
		Flavors: s.Flavors,
		Score: s.Score,
		Price: s.Price,
		Link: s.Link,
		Favorite: s.Favorite,
	}
}
