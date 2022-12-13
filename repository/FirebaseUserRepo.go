package repository

import (
	"app/domain"
	//"app/presenter"
	"context"
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

type UserFirebase interface{
	CreateUser(context.Context, domain.User)(domain.User, error)
	ReadID(context.Context, string, domain.User)(domain.User, error)
}

type Firebase struct{
	client	firestore.Client
}

func NewFirebaseRepository(client firestore.Client) *Firebase{
	return &Firebase{client : client}//retorna o valor para struct Firebase --> com valor atribuido na main
}

func (f Firebase) EmailRegistered(c context.Context, email string) (bool) {
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

func (f Firebase) CreateUser(c context.Context, u domain.User) (user domain.User, err error){
	usersCollection := f.client.Collection("Users")
	
	if f.EmailRegistered(c, u.Email) == false{
			u.ID = uuid.NewString()
			//service.UserUseCase.Create(c,u)
		_, err= usersCollection.Doc(u.ID).Create(c, u)
		if err != nil{
			return u, err
			}
		return u, nil
		}
	return u, nil
}

//context errado?
func (f Firebase) ReadID(c context.Context, id string, u domain.User) (user domain.User, err error){
	usersCollection := f.client.Collection("Users")

	doc, err := usersCollection.Doc(id).Get(c)//talvez GET esteja errado
	if err != nil {
		return user, err
	}
	if err := doc.DataTo(&u); err != nil {
			return u, err
	}
	return u, nil
}