package main

import (
	"investPedia/auth"
	"investPedia/handler"
	"investPedia/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/investpedia2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// getting env variables SITE_TITLE and DB_HOST
	secretKeyJWT := os.Getenv("SECRET_KEY_JWT")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService(secretKeyJWT)
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/checkEmail", userHandler.CheckEmailAvailibility)
	api.POST("/uploadAvatar", userHandler.UploadAvatar)
	router.Run()

}
