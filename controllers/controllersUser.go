package controllers

import (
	"app/domain"
	//"app/repository"
	//"app/presenter"
	"net/http"
	"errors"
	"context"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

type UserControllers interface {
/* 	CreateUser(c *gin.Context)
	GetIDUser(c *gin.Context) */
	CreateUser(context.Context, domain.User)(domain.User, error)
	ReadID(context.Context, string, domain.User)(domain.User, error)
}

type UserControllerService struct{
	repository UserControllers
}

func NewUserControllerService (repository UserControllers) *UserControllerService{
	return &UserControllerService{repository: repository}
}

func (uc UserControllerService) CreateUser(c *gin.Context) {
	var u domain.User
	//var user presenter.User
	var ve validator.ValidationErrors
	
	if err := c.ShouldBindJSON(&u); err != nil {
		if errors.As(err, &ve) {
			out := make([]domain.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = domain.ErrorMsg{Field: fe.Field(), Message: domain.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errors": out})
		}
		return
	}
	user, err := uc.repository.CreateUser(c ,u) //passar para present --> present.user, err := repository.Create(c, u)
	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro ao criar dado na struct domain.User": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func  (uc UserControllerService) GetIDUser(c *gin.Context) {
	var u domain.User
	//var user presenter.User

	givenID := c.Params.ByName("id")
	user, err := uc.repository.ReadID(c, givenID, u)//desta forma que passa a função?
	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro ao extrair dado da struct User": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
