package repository

import (
	"context"
	"app/domain"
//	"log"
	//"net/http"
	//"os"
	//"strconv"
	//"cloud.google.com/go/firestore"

	"google.golang.org/api/iterator"
	//"github.com/google/uuid"
	//"errors"
	//"github.com/go-playground/validator/v10"
	"fmt"
	//"sort"
)

func nameRegistered(f Firebase, c context.Context, name string) bool{
	shopCollection := f.client.Collection("Shops")

	iter := shopCollection.Where("name", "==", name).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			fmt.Println(err)
		}
		if doc != nil{
			fmt.Println("Nome da loja já foi cadastrada.")
			return true
		}
	}	
	fmt.Println("Nome da loja pode ser criada")
		return false
}	

func (f Firebase) CreateShop (c context.Context, s domain.Shop) (domain.Shop, error){
	shopCollection := f.client.Collection("Shops")

	if nameRegistered(f, c, s.Name) == false{
		_, err := shopCollection.Doc(s.ID).Create(c, s)
		if err != nil {
			fmt.Println(err)
			return s, err
		}
		return s, err
	}
	return domain.Shop{}, nil
}

/* func GetAll(c context.Context, s domain.Shop) (domain.Shop, error) {

}
func GetByID(c context.Context, id string, s domain.Shop) (domain.Shop, error){

}

func GetByName(c context.Context, name string, s domain.Shop) (domain.Shop, error){

}

func  GetByScore(c context.Context, score string, s domain.Shop) (domain.Shop, error){

}
	
func GetByPrice(c context.Context, price string, s domain.Shop) (domain.Shop, error){

}
	
func Update(c context.Context, s domain.Shop) (domain.Shop, error){

}
	
func UpdateField(c context.Context, score string, s domain.Shop) (domain.Shop, error){

}
	
func Delete(c context.Context, id string, s domain.Shop) (domain.Shop, error){
	
} */