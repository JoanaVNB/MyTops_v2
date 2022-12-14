package repository

import (
	"app/domain"
	//"app/presenter"
	//"app/service"
	"context"
	"cloud.google.com/go/firestore"
	"fmt"
	//"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

type DB interface{
	Create(context.Context, domain.User) (domain.User, error)
	GetID(context.Context, string, domain.User) (domain.User, error)
}

type Firebase struct{
	client	firestore.Client
}

func NewFirebaseRepository(client firestore.Client) *Firebase{
	return &Firebase{client : client}
}

func EmailRegistered(f Firebase, c context.Context, email string) (bool) {
	usersCollection := f.client.Collection("Users")

	iter := usersCollection.Where("email", "==", email).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			fmt.Println(err)
		}
		if doc != nil{
			fmt.Println("E-mail foi cadastrado.")
			return true
		}
	}
	fmt.Println("Conta pode ser criada")
	return false
}	

func (f Firebase) Create(c context.Context, u domain.User) (user domain.User, err error){
	usersCollection := f.client.Collection("Users")
	
	if EmailRegistered(f, c, u.Email) == false{
			//u.ID = uuid.NewString()
			//service.UserRepository.Create(c, u) PQ NÃO FUNCIONA?
			_, err= usersCollection.Doc(u.ID).Create(c, u)
		if err != nil{
			return u, err
			}
		return u, nil
		}
	return u, nil
}

func (f Firebase) GetID(c context.Context, id string, u domain.User) (user domain.User, err error){
	usersCollection := f.client.Collection("Users")

	doc, err := usersCollection.Doc(id).Get(c)
	if err != nil {
		return user, err
	}
	if err := doc.DataTo(&u); err != nil {
			return u, err
	}
	return u, nil
}