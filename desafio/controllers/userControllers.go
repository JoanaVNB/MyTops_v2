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

type UserRepository interface {
	Create(c context.Context, u domain.User) (domain.User, error)
	GetID(context.Context, string, domain.User) (domain.User, error)
	Login(context.Context, domain.User, domain.Login) bool
} 

type UserController struct{
	repository UserRepository
}

func NewUserController(repository UserRepository) *UserController{
	return &UserController{repository: repository}
}

func (uc UserController) Create(c *gin.Context) {
	var u domain.User
	var ve validator.ValidationErrors
	
	if err := c.ShouldBindJSON(&u); err != nil {
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
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
	if user.Name == ""{//o user não foi cadastrado caso o campo retornar vazio, é pq não foi cadastrado
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
	
	bol := uc.repository.Login(c, u, l)
	if bol{
		c.JSON(http.StatusAccepted,  "Usuário autorizado")
	}
	if !bol{
		c.JSON(http.StatusBadRequest, "Senha incorreta")	
	}
}
