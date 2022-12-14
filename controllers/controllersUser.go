package controllers

import (
	"app/domain"
	"app/repository"
	//"app/service"
	//"app/repository"
	//"app/presenter"
	"errors"
	"net/http"

	//"context"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

type UserController struct{
	repository repository.UserRepository
}

func NewUserController(repository repository.UserRepository) *UserController{
	return &UserController{repository: repository}
}

func (uc UserController) Create(c *gin.Context) {
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
	user, err := uc.repository.Create(c, u) //passar para present --> present.user, err := repository.Create(c, u)
	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro ao criar dado na struct domain.User": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func  (uc UserController) GetID(c *gin.Context) {
	var u domain.User
	//var user presenter.User

	givenID := c.Params.ByName("id")
	user, err := uc.repository.GetID(c, givenID, u)
	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro ao extrair dado da struct User": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
