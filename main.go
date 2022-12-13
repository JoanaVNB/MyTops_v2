package main

import(
	"os"
	"context"
	"cloud.google.com/go/firestore"
	"log"
	"app/repository"
	"app/controllers"
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
	//userService := service.NewUserUseCase(firebaseRepository)// atribui a qual repo os services devem se conectar
	userControllers := controllers.NewUserControllerService(firebaseRepository)

//Conectando ao Gin
	r := gin.Default()
	
		r.POST("/user", userControllers.CreateUser)
		r.GET("/user/:id", userControllers.GetIDUser)
	
		r.Run(":5500")
}
