package controllers

import (
	"app/domain"
	"app/presenters"
	"app/service"
	"errors"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

type ShopController struct{
	repository service.ShopRepository
}

func NewShopController(repository service.ShopRepository) *ShopController{
	return &ShopController{repository: repository}
}

func (sc ShopController) CreateShop(c *gin.Context){
	var s domain.Shop
	var ve validator.ValidationErrors

	if err := c.ShouldBindJSON(&s); err != nil {
		if errors.As(err, &ve){
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve{
				out[i] = ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
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
	if shop.Name == ""{
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

func (sc ShopController) Update(c *gin.Context){
	var s domain.Shop
	var ve validator.ValidationErrors
	givenID := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&s); err != nil {
		if errors.As(err, &ve){
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve{
				out[i] = ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": out})
		}
		return 
	}

	_, err := sc.repository.Update(c, givenID, s)
	if err != nil{
		c.JSON(http.StatusBadRequest, "não encontrado")
		return
	}
	c.JSON(http.StatusOK, "atualizado")
}

func (sc ShopController) UpdateScore(c *gin.Context){
	var s domain.Shop
	givenID := c.Params.ByName("id")
	score, _ := strconv.ParseFloat(c.Param("score"), 64)

	_, err := sc.repository.UpdateScore(c, givenID, score, s)
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

	_, err := sc.repository.UpdatePrice(c, givenID, p.NewPrice, s)
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
	_, keys := sc.repository.ListScores(c, s)
	
	c.JSON(http.StatusOK, keys)
}


