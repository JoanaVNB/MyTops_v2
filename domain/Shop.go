package domain

type Shop struct {
	ID      string  `firestore:"id"`
	Name    string  `firestore:"name" binding:"required"`
	Flavors [6]string `firestore:"flavors"`
	Score   float64 `firestore:"score" binding:"required"`
	Price   float64 `firestore:"price"`
	Link    string  `firestore:"link"`
	Favorite	bool  `firestore:"favorite"`
}

type NameUpdated struct{
	NewName	string	`firestore:"newname"`
}

type PriceUpdated struct{
	NewPrice	float64	`firestore:"newprice"`
}


