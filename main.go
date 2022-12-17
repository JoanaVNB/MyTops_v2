package main

import(
	"os"
	"context"
	"cloud.google.com/go/firestore"
	"log"
	"app/repository/Firebase"
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
	shopService := service.NewShopService(firebaseRepository)
	shopControllers := controllers.NewShopController(shopService)

//Conectando ao Gin
	r := gin.Default()

//User
		r.POST("/users", userControllers.Create)
		r.GET("/users/:id", userControllers.GetID)
		r.POST("/login", userControllers.Login)

//Shop
		r.POST("/pizzerias", shopControllers.CreateShop)
		r.GET("/pizzerias", shopControllers.GetAll)
		r.GET("/pizzerias/:id", shopControllers.GetByID)
		r.GET("/pizzerias/name/:name", shopControllers.GetByName)
		r.GET("/pizzerias/score/:score", shopControllers.GetByScore)
		r.GET("/pizzerias/price/:price", shopControllers.GetByPrice)
		r.GET("/pizzerias/score/:score/price/:price", shopControllers.GetByScorePrice)//não funciona -> como utilizo & e um alias?
		r.PUT("/pizzerias/:id", shopControllers.Update)
		r.PUT("/pizzerias/:id/score/:score", shopControllers.UpdateScore)
		r.PUT("/pizzerias/:id/price", shopControllers.UpdatePrice)//errado
		r.DELETE("/pizzerias/:id", shopControllers.Delete)
		r.GET("/pizzerias/ranking", shopControllers.Ranking)
		
	r.Run(":5500")
}
