package repository

import (
	"context"
	"app/domain"
//	"log"
	//"net/http"
	//"os"
	//"strconv"
	"cloud.google.com/go/firestore"

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

func (f Firebase) GetAll(c context.Context, s domain.Shop) ([]domain.Shop, error) {
	shopCollection := f.client.Collection("Shops")
	var docs = make([]domain.Shop,0)

	iter := shopCollection.Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			return docs, err
		}
		if err := doc.DataTo(&s); err != nil {
		return docs, err
		}
	docs = append(docs, s)	
	}
	return docs, nil
}


func (f Firebase) GetByID(c context.Context, id string, s domain.Shop) (domain.Shop, error){
	shopCollection := f.client.Collection("Shops")

	doc, err := shopCollection.Doc(id).Get(c)
	if err != nil {
		return domain.Shop{}, err
	}

	if err := doc.DataTo(&s); err != nil {
		return s, err
	}
	s.ID = doc.Ref.ID
	return s, nil
}

func (f Firebase) GetByName(c context.Context, name string, s domain.Shop) (domain.Shop, error){
	shopCollection := f.client.Collection("Shops")
	var docum  *firestore.DocumentSnapshot

	iter := shopCollection.Where("name", "==", name).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			return domain.Shop{}, err
		}
		docum = doc
	}		
	if err := docum.DataTo(&s); err != nil {
		return domain.Shop{}, err
	}
	return s, nil
}

func (f Firebase) GetByScore(c context.Context, score float64, s domain.Shop) ([]domain.Shop, error){
	shopCollection := f.client.Collection("Shops")
	var docs = make([]domain.Shop,0)

	iter := shopCollection.Where("score", ">=", score).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			return docs, err
		}
		if err := doc.DataTo(&s); err != nil {
			return docs, err
		}
		docs = append(docs, s)
	}		
	return docs, nil
}

func (f Firebase) GetByPrice(c context.Context, price float64, s domain.Shop) ([]domain.Shop, error){
	shopCollection := f.client.Collection("Shops")
	var docs = make([]domain.Shop,0)

	iter := shopCollection.Where("price", "<=", price).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			return docs, err
		}
		if err := doc.DataTo(&s); err != nil {
			return docs, err
		}
		docs = append(docs, s)
	}		
	return docs, nil
}

func (f Firebase) GetByScorePrice(c context.Context, score float64, price float64, s domain.Shop) ([]domain.Shop, error){
	var docs = make([]domain.Shop,0)

	shops, err :=f.GetByScore(c, score, s); if err !=nil {
		return docs, err
	}
	for _, value := range shops{
		eachShopByScore := value
		shopsByPrice, err := f.GetByPrice(c, score, eachShopByScore); if err != nil {
			return docs, err}
		for _, eachShopByPrice := range shopsByPrice{
			docs = append(docs, eachShopByPrice)
		}
	}
	return docs, nil
}


func (f Firebase) Update(c context.Context, id string, s domain.Shop) (error){
	shopCollection := f.client.Collection("Shops")

	doc, err := shopCollection.Doc(id).Get(c)
	if err != nil{
		return err
	}
	/* s.ID = uuid.NewString() */
	s.ID = doc.Ref.ID
	_, err = shopCollection.Doc(id).Set(c, s)
	if err != nil{
		return  err
	}
	return nil
}
	
func (f Firebase) UpdateScore(c context.Context, id string, score float64, s domain.Shop) (error){
	shopCollection := f.client.Collection("Shops")

	_, err := shopCollection.Doc(id).Update(c, []firestore.Update{{Path: "score", Value: score}})
	if err !=nil {
		return err
	}
	return nil
}

func (f Firebase) UpdatePrice(c context.Context, id string, price float64, s domain.Shop) (error){
	shopCollection := f.client.Collection("Shops")
	var p domain.PriceUpdated
	
	_, err := shopCollection.Doc(id).Update(c, []firestore.Update{{Path: "price", Value: p.NewPrice}})
	if err !=nil {
		return err
	}
	return nil
}
	
func (f Firebase) Delete(c context.Context, id string) ( error){
	shopCollection := f.client.Collection("Shops")

	_, err := shopCollection.Doc(id).Delete(c); if err != nil{
		return  err
	}
	return nil
}

func (f Firebase) ListScores(c context.Context, s domain.Shop) (map[string]float64){
	shopCollection := f.client.Collection("Shops")
	scores := make(map[string]float64)

	iter := shopCollection.Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			return nil
		}
		if err := doc.DataTo(&s); err != nil{
			return nil
		}
		scores[s.Name] = s.Score
	}
	return scores
}


	

