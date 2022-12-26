package controllers

import (
	"app/domain"
	"app/presenters"
	repository "app/repository/Firebase"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"app/service"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
)

func connectRepo() (*UserController, *gin.Engine){
	_ = os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:9090")

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "desafio-c0479")
	if err != nil {
		log.Println(err)
	}
	repo := repository.NewFirebaseRepository(*client)
	userService := service.NewUserService(repo)
	userControllers := NewUserController(userService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return userControllers, r
}

func Test_Create_Code201(t *testing.T){
	//conecta banco de dados e gin

	userControllers, r := connectRepo()
	r.POST("/users", userControllers.Create)

	//request & response

	u := domain.User{
			Name: "Joana",
			Email: "joanavidon@gmail.com",
			Password: "jo123"}
	
	uJSON, err := json.Marshal(u)
	if err != nil{
		t.Fatal(err)
	}

	var pu presenters.User
		
	path:= "/users"
	req, _ := http.NewRequest("POST", path,
				bytes.NewBuffer(uJSON))
	response := httptest.NewRecorder()	

	r.ServeHTTP(response, req)

	err = json.Unmarshal(response.Body.Bytes(), &pu); if err != nil{
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.NotEmpty(t, pu.ID)
	assert.Equal(t, "Joana", pu.Name)
	assert.Equal(t, "joanavidon@gmail.com", pu.Email)
}

func Test_Create_Code400(t *testing.T){
	//conecta banco de dados e gin

	userControllers, r := connectRepo()
	r.POST("/users", userControllers.Create)

	//request & response

		u := domain.User{Name: "Joana",
				Email: "joanavidon@gmail.com",
				Password: "jo123"}
		uJSON, _ := json.Marshal(u)
		
	path:= "/users"
	req, _ := http.NewRequest("POST", path,
				bytes.NewBuffer(uJSON))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

