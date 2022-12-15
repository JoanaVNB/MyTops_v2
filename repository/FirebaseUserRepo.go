package repository

import (
	"app/domain"
	"context"
	"fmt"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

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
			fmt.Println("Este e-mail já foi cadastrado.")
			return true
		}
	}
	fmt.Println("Conta pode ser criada.")
	return false
}	

func (f Firebase) Create(c context.Context, u domain.User) (domain.User, error){
	usersCollection := f.client.Collection("Users")
	
	if EmailRegistered(f, c, u.Email) == false{
			_, err := usersCollection.Doc(u.ID).Create(c, u)
			if err != nil{
				return u, err
			}
		return u, nil
	}
	return domain.User{}, nil
}

func (f Firebase) GetID(c context.Context, id string, u domain.User) (domain.User, error){
	usersCollection := f.client.Collection("Users")

	doc, err := usersCollection.Doc(id).Get(c)
	if err != nil {
		return u, err
	}
	if err := doc.DataTo(&u); err != nil {
			return u, err
	}
	return u, nil
}

func (f Firebase) Login(c context.Context, u domain.User, l domain.Login) bool{
	usersCollection := f.client.Collection("Users")	

	iter := usersCollection.Where("email", "==", l.Email).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			fmt.Println(err)
		}
	
		if err := doc.DataTo(&u); err != nil {
			fmt.Println(err)
		}
	}

	if u.Password != l.Password{
		return false
	}
	return true
}