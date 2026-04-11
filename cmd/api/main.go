package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RichMagg/gurizes-economy-service/cmd/internal/controllers"
	"github.com/RichMagg/gurizes-economy-service/cmd/internal/database"
	"github.com/RichMagg/gurizes-economy-service/cmd/internal/repositories"
	"github.com/RichMagg/gurizes-economy-service/cmd/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(".env not found!")
	}

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	if connString == "" {
		panic("POSTGRES variables is not set!")
	}

	db := database.NewDB(connString)

	server := gin.Default()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("Pong!"))
	})

	server.GET("/users", userController.GetUsers)
	server.POST("/users", userController.CreateUser)

	server.Run(":8080")
}
