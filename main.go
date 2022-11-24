package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/templateOfService/connectors/mysql"
	_ "github.com/templateOfService/docs"
	"github.com/templateOfService/services/auth"
	"github.com/templateOfService/services/core"
	"log"
	"os"
)

func initRouter() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()
	authHandler := auth.NewHandler()
	router.POST("/api/v1/auth/signup", authHandler.Signup)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	coreHandler := core.NewHandler()
	router.GET("/api/v1/words/search", coreHandler.Search)
	router.POST("/api/v1/words/add", coreHandler.AddWord)
	return router
}

// @title User API documentation
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env: %s", err.Error())
	}

	err = mysql.Connect()
	if err != nil {
		log.Fatalf("Error connecting to mysql: %s", err.Error())
	}

	router := initRouter()
	err = router.Run()
	if err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}
