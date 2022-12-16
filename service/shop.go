package service

import (
	"app/domain"
	"context"
	"github.com/google/uuid"
)

type ShopRepository interface {
	CreateShop(context.Context, domain.Shop) (domain.Shop, error)
	// GetAll(context.Context, domain.Shop) (domain.Shop, error)
	// GetByID(context.Context, string, domain.Shop) (domain.Shop, error)
	// GetByName(context.Context, string, domain.Shop) (domain.Shop, error)
	// GetByScore(context.Context, string, domain.Shop) (domain.Shop, error)
	// GetByPrice(context.Context, string, domain.Shop) (domain.Shop, error)
	// Update(context.Context, domain.Shop) (domain.Shop, error)
	// UpdateField(context.Context, string, domain.Shop) (domain.Shop, error)
	// Delete(context.Context, string, domain.Shop) (domain.Shop, error)
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

