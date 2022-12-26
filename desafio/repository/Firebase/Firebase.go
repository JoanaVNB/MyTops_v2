package repository

import (
	"cloud.google.com/go/firestore"

)

type Firebase struct{
	client	firestore.Client
}

func NewFirebaseRepository(client firestore.Client) *Firebase{
	return &Firebase{client : client}
}