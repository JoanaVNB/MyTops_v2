package presenters

import (
	"app/domain"
)

type Shop struct {
	Name    string  `json:"name"`
	Flavors [6]string `json:"flavors"`
	Score   float64  `json:"score"`
	Price   float64 `json:"price"`
	Link    string  `json:"link"`
	Favorite	bool `json:"favorite"`
}

type PriceUpdated struct{
	NewPrice	float64 `json:"newprice"`
}

func PresenterShop(s domain.Shop) *Shop{
	return &Shop{
		Name: s.Name,
		Flavors: s.Flavors,
		Score: s.Score,
		Price: s.Price,
		Link: s.Link,
		Favorite: s.Favorite,
	}
}
