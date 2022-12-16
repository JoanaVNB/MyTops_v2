package controllers

import (
	"app/domain"
	"context"
	"app/presenters"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

type ShopRepository interface {
	CreateShop(context.Context, domain.Shop) (domain.Shop, error)
// 	GetAll(context.Context, domain.Shop) (domain.Shop, error)
// 	GetByID(context.Context, string, domain.Shop) (domain.Shop, error)
// 	GetByName(context.Context, string, domain.Shop) (domain.Shop, error)
// 	GetByScore(context.Context, string, domain.Shop) (domain.Shop, error)
// 	GetByPrice(context.Context, string, domain.Shop) (domain.Shop, error)
// 	Update(context.Context, domain.Shop) (domain.Shop, error)
// 	UpdateField(context.Context, string, domain.Shop) (domain.Shop, error)
// 	Delete(context.Context, string, domain.Shop) (domain.Shop, error)
} 

type ShopController struct{
	repository ShopRepository
}

func NewShopController(repository ShopRepository) *ShopController{
	return &ShopController{repository: repository}
}

func (sc ShopController) CreateShop(c *gin.Context){
	var s domain.Shop
	var ve validator.ValidationErrors

	if err := c.ShouldBindJSON(&s); err != nil {
		if errors.As(err, &ve){
			out := make([]domain.ErrorMsg, len(ve))
			for i, fe := range ve{
				out[i] = domain.ErrorMsg{Field: fe.Field(), Message: domain.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": out})
		}
		return 
	}

	shop, err := sc.repository.CreateShop(c, s)
	presenterShop := *presenters.PresenterShop(shop)
	if  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao criar dado na struct domain.Shop": err.Error()})
		return
	}
	if shop.Name == ""{//se campo retornou vazio, é pq não foi cadastrado
		c.JSON(http.StatusBadRequest, "Loja não foi cadastrada.")
		return
	}
	c.JSON(http.StatusCreated, presenterShop)
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