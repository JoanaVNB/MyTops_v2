package service

import (
	"app/domain"
	"context"
	"github.com/google/uuid"
)

type ShopRepository interface {
	CreateShop(context.Context, domain.Shop) (domain.Shop, error)
	GetAll(context.Context, domain.Shop) ([]domain.Shop, error)
	GetByID(context.Context, string, domain.Shop) (domain.Shop, error)
	GetByName(context.Context, string, domain.Shop) (domain.Shop, error)
	GetByScore(context.Context, float64, domain.Shop) ([]domain.Shop, error)
	GetByPrice(context.Context, float64, domain.Shop) ([]domain.Shop, error)
	GetByScorePrice(context.Context, float64, float64, domain.Shop) ([]domain.Shop, error)
	Update(context.Context, string, domain.Shop) (error)
	UpdateScore(context.Context, string, float64, domain.Shop) (error)
	UpdatePrice(context.Context, string, float64, domain.Shop) (error)
	Delete(context.Context, string) (error)
	ListScores(context.Context, domain.Shop) (map[string]float64)
} 

type ShopService struct {
	repository ShopRepository
}

func NewShopService (repository ShopRepository) *ShopService{
	return &ShopService{repository: repository}
}

//COMO É VAZIO PARA INT?
func validateShop(s domain.Shop) bool{
	if s.Name == "" || s.Score == 0 || s.Price == 0 {
		return false
	}
	return true
}

func (ss ShopService) CreateShop(c context.Context, s domain.Shop) (domain.Shop, error){
	if validateShop(s) == true {
		s.ID = uuid.NewString()
		shop, err := ss.repository.CreateShop(c,s)
		shop.ID = s.ID
		if err != nil{
			return shop, err
		}
		return shop, nil
	}
	return domain.Shop{}, nil
}

func (ss ShopService)  GetAll(c context.Context, s domain.Shop) ([]domain.Shop, error) {
	shop, err := ss.repository.GetAll(c, s)
	if err != nil{
		return  shop, err
	}
	return shop, nil
}

func (ss ShopService) GetByID(c context.Context, id string, s domain.Shop) (domain.Shop, error){
	shop, err := ss.repository.GetByID(c, id, s)
	shop.ID = s.ID
	if err != nil{
		return shop, err
	}
	return shop, nil
}

func (ss ShopService) GetByName(c context.Context, name string, s domain.Shop) (domain.Shop, error){
	shop, err := ss.repository.GetByName(c, name, s)
	if err != nil{
		return  domain.Shop{}, err
	}
	return shop, nil
}

func (ss ShopService) GetByScore(c context.Context, score  float64, s domain.Shop) ([]domain.Shop, error){
	shops, err := ss.repository.GetByScore(c, score, s)
	if err != nil{
		return  shops, err
	}
	return shops, nil
}
	
func (ss ShopService) GetByPrice(c context.Context, price float64, s domain.Shop) ([]domain.Shop, error){
	shops, err := ss.repository.GetByPrice(c, price, s)
	if err != nil{
		return  shops, err
	}
	return shops, nil
}

func (ss ShopService) GetByScorePrice(c context.Context, score float64, price float64, s domain.Shop) ([]domain.Shop, error){
	shops, err := ss.repository.GetByScorePrice(c, score, price, s)
	if err != nil{
		return  shops, err
	}
	return shops, nil
}

func (ss ShopService) Update(c context.Context, id string, s domain.Shop) (error){
	err := ss.repository.Update(c, id, s)
	if err != nil{
		return  err
	}
	return nil
}
	
func (ss ShopService) UpdateScore(c context.Context, id string, score float64, s domain.Shop) (error){
	err := ss.repository.UpdateScore(c, id, score, s)
	if err != nil{
		return err
	}
	return nil
}

func (ss ShopService) UpdatePrice(c context.Context, id string, price float64, s domain.Shop) (error){
	err := ss.repository.UpdatePrice(c, id, price, s)
	if err != nil{
		return err
	}
	return nil
}
	
func (ss ShopService) Delete(c context.Context, id string) (error){
	return ss.repository.Delete(c, id)
}

func (ss ShopService) ListScores(c context.Context, s domain.Shop) (map[string]float64){
	return ss.repository.ListScores(c, s)
}