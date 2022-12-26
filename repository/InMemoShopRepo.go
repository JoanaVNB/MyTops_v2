package repository

import (
	"app/domain"
	"context"
	"fmt"
)

type InMemoShopRepository struct{
	sMap	map[string]domain.Shop
}

func NewInMemoShopRepository() *InMemoShopRepository{
	return &InMemoShopRepository{sMap: make(map[string]domain.Shop)}
}

func (in InMemoShopRepository) CreateShop(c context.Context, s domain.Shop) (domain.Shop, error){
	in.sMap[s.ID] = s
	return s,  nil
}

func (in InMemoShopRepository) GetAll(c context.Context, s domain.Shop) ([]domain.Shop, error){
	var shops []domain.Shop
	for _, k := range in.sMap{
		shop := k
		shops = append(shops, shop)
	}	
	return shops, nil
}

func (in InMemoShopRepository) GetByID(c context.Context, id string, s domain.Shop) (domain.Shop, error){
	for _, k := range in.sMap{
		gotID := k.ID
		if gotID == id{
			return s, nil
		}
	}
	return s, nil
}

func (in InMemoShopRepository) GetByName(c context.Context, name string, s domain.Shop) (domain.Shop, error){
for _, k := range in.sMap{
	gotName := k.Name
	if gotName == name{
		return s, nil
	}
}
return s, nil
}

func (in InMemoShopRepository) GetByScore(c context.Context, score float64, s domain.Shop) ([]domain.Shop, error){
	var shops []domain.Shop
	for _, k := range in.sMap{
		gotScore := k.Score
		if gotScore == score{
			shops = append(shops, s)
		}
	}
	return shops, nil
}
	
func (in InMemoShopRepository) GetByPrice(c context.Context, price float64, s domain.Shop) ([]domain.Shop, error){
	var shops []domain.Shop
	for _, k := range in.sMap{
		gotPrice := k.Price
		if gotPrice == price{
			shops = append(shops, s)
		}
	}
	return shops, nil
}

func (in InMemoShopRepository) Update(c context.Context, id string, s domain.Shop) (error){
	in.sMap[id] = s
	return nil
}

func (in InMemoShopRepository) UpdateScore(c context.Context, id string, score float64, s domain.Shop) (error){
	s = domain.Shop{Score: score}
	in.sMap[id] = s
	return nil
}

func (in InMemoShopRepository) UpdatePrice(c context.Context, id string, price float64, s domain.Shop) (error){
	s = domain.Shop{Price: price}
	in.sMap[id] = s
	return nil
}

func (in InMemoShopRepository) Delete(d context.Context, id string) (error){
	delete(in.sMap, id)
	return nil
}

func (in InMemoShopRepository) ListScores(c context.Context, s domain.Shop) (map[string]float64, []string){
	scores := make(map[string]float64)

	shops, err := in.GetAll(c, s)
	for _, shop := range shops{ 
		scores[shop.Name] = shop.Score
	}
	if err != nil{
		return nil, nil
	}
	return scores, nil
}

