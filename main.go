package main

import (
	"context"
	"log"
	"net/http"
	config "pioApi/config"
	controllers "pioApi/controllers"
	_ "pioApi/docs"
	"pioApi/repositories"
	"pioApi/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	if err := config.DB.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		defer config.DB.Close()
	}

	uc := controllers.NewUserController()
	us := services.NewUserService()
	ur := repositories.NewUserRepository()
	us.SetRepository(ur)
	uc.SetService(us)
	uc.InitRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

//	@title			Pio game api
//	@version		1.0
//	@description	Pio game rewrtitten as api

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
