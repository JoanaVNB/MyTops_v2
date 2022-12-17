package controllers

import (
	"app/domain"
	"app/presenters"
	"context"
	"errors"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"sort"
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
	ListScores(c context.Context, s domain.Shop) (map[string]float64)
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

func (sc ShopController) GetAll(c *gin.Context){
	var s domain.Shop

	shops, err := sc.repository.GetAll(c,s)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		return
	}

	for _, sh := range shops{
		shop := sh
		presenterShop := *presenters.PresenterShop(shop)
		c.JSON(http.StatusOK, presenterShop)
	}
}

func (sc ShopController) GetByID(c *gin.Context){
	var s domain.Shop

	givenID, _:= c.Params.Get("id")
	shop, err := sc.repository.GetByID(c, givenID, s)
	if err != nil{
		c.JSON(http.StatusBadRequest, "não encontrado")
		return
	}
	presenterShop := *presenters.PresenterShop(shop)
	c.JSON(http.StatusOK, presenterShop)
}

func (sc ShopController) GetByName(c *gin.Context){
	var s domain.Shop

	name:= c.Params.ByName("name")
	shop, err := sc.repository.GetByName(c, name, s)
	if err != nil{
		c.JSON(http.StatusBadRequest, "não encontrado")
		return
	}
	presenterShop := *presenters.PresenterShop(shop)
	c.JSON(http.StatusOK, presenterShop)
}

func (sc ShopController) GetByScore(c *gin.Context){
	var s domain.Shop

	score, err:= strconv.ParseFloat(c.Param("score"), 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, "erro ao converter para float64")
	}

	shops, err := sc.repository.GetByScore(c, score, s)
	if err != nil{
		c.JSON(http.StatusBadRequest, "não encontrado")
		return
	}

	for _, sh := range shops{
		shop := sh
		presenterShop := *presenters.PresenterShop(shop)
		c.JSON(http.StatusOK, presenterShop)
	}
}

func (sc ShopController) GetByPrice(c *gin.Context){
	var s domain.Shop

	price, err:= strconv.ParseFloat(c.Param("price"), 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, "erro ao converter para float64")
	}

	shops, err := sc.repository.GetByPrice(c, price, s)
	if err != nil{
		c.JSON(http.StatusBadRequest, "não encontrado")
		return
	}

	for _, sh := range shops{
		shop := sh
		presenterShop := *presenters.PresenterShop(shop)
		c.JSON(http.StatusOK, presenterShop)
	}
}

func (sc ShopController) GetByScorePrice(c *gin.Context){
	var s domain.Shop

	score, err:= strconv.ParseFloat(c.Param("score"), 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, "erro ao converter para float64")
	}

	price, err:= strconv.ParseFloat(c.Param("price"), 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, "erro ao converter para float64")
	}

	shops, err := sc.repository.GetByScorePrice(c, score, price, s)
	if err != nil{
		c.JSON(http.StatusBadRequest, "não encontrado")
		return
	}

	for _, sh := range shops{
		shop := sh
		presenterShop := *presenters.PresenterShop(shop)
		c.JSON(http.StatusOK, presenterShop)
	}
}

func (sc ShopController) Update(c *gin.Context){
	var s domain.Shop
	var ve validator.ValidationErrors
	givenID := c.Params.ByName("id")

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

	err := sc.repository.Update(c, givenID, s)
	if err != nil{
		c.JSON(http.StatusBadRequest, "não encontrado")
		return
	}
	c.JSON(http.StatusOK, "atualizado")
}

func (sc ShopController) UpdateScore(c *gin.Context){
	var s domain.Shop
	givenID := c.Params.ByName("id")
	score, err:= strconv.ParseFloat(c.Param("score"), 64)

	err = sc.repository.UpdateScore(c, givenID, score, s)
	if err != nil{
		c.JSON(http.StatusBadRequest, "não encontrado")
		return
	}
	c.JSON(http.StatusOK, "atualizado")
}

func (sc ShopController) UpdatePrice(c *gin.Context){
	var s domain.Shop
	var p domain.PriceUpdated
	givenID := c.Params.ByName("id")
	
	if err := c.ShouldBindJSON(&p); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao converter": err.Error()})
		return
	}

	err := sc.repository.UpdatePrice(c, givenID, p.NewPrice, s)
	if err != nil{
		c.JSON(http.StatusBadRequest, "não encontrado")
		return
	}
	c.JSON(http.StatusOK, "atualizado")
}

func (sc ShopController) Delete(c *gin.Context){

	givenID := c.Params.ByName("id")
	err := sc.repository.Delete(c, givenID)
	if err != nil{
		c.JSON(http.StatusBadRequest, "não encontrado")
		return
	}
	c.JSON(http.StatusOK, "deletado")
}

func (sc ShopController) Ranking(c *gin.Context){
	var s domain.Shop
	list := sc.repository.ListScores(c, s)
	
	keys := make([]string, 0, len(list))

	for k := range list{
		keys = append(keys, k)
	}
	
	sort.SliceStable(keys, func(i, j int) bool{
		return list[keys[i]] > list[keys[j]]
	})

	c.JSON(http.StatusOK, keys)
}


