package controllers

import (
	"app/domain"
	"app/repository"
	"app/presenters"
	"errors"
	"net/http"
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
	user, err := uc.repository.Create(c, u)
	presenterUser := *presenters.PresenterUser(user)
	if  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao criar dado na struct domain.User": err.Error()})
		return
	}
	if user.Name == ""{//se campo retornou vazio, é pq não foi cadastrado
		c.JSON(http.StatusBadRequest, "Usuário não foi cadastrado.")
		return
	}
	c.JSON(http.StatusCreated, presenterUser)
}

func  (uc UserController) GetID(c *gin.Context) {
	var u domain.User

	givenID := c.Params.ByName("id")
	user, err := uc.repository.GetID(c, givenID, u)
	if  err != nil {
		c.JSON(http.StatusBadRequest, "ID não encontrado")
		return
	}
	presenterUser := *presenters.PresenterUser(user)
	c.JSON(http.StatusOK, &presenterUser)
}

func (uc UserController) Login(c *gin.Context){
	var u domain.User
	var l domain.Login

	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro ao extrair dado da struct Login": err.Error()})
			return
	}
	
	bol, _ := uc.repository.Login(c, u, l)
	if bol == true{
		c.JSON(http.StatusAccepted,  "Usuário autorizado")
	}
	if bol == false{
		c.JSON(http.StatusBadRequest, "Senha incorreta")	
	}
}
