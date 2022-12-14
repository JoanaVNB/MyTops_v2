package main

import(
	"os"
	"context"
	"cloud.google.com/go/firestore"
	"log"
	"app/repository"
	"app/controllers"
	"app/service"
	"github.com/gin-gonic/gin"
)

func main(){

//Conectando User ao Firebase
	_ = os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:9090")

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "desafio-c0479")
	if err != nil {
		log.Println(err)
	}

	firebaseRepository := repository.NewFirebaseRepository(*client) //atribui o client ao Repository --> inicia o doc no repositório
	
//Services
	userService := service.NewUserService(firebaseRepository)// atribui a qual repo os services devem se conectar
	userControllers := controllers.NewUserController(userService) //injetar o userService

//Conectando ao Gin
	r := gin.Default()
	
		r.POST("/user", userControllers.Create)
		r.GET("/user/:id", userControllers.GetID)
	
	r.Run(":5500")
}
