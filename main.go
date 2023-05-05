package main

import (
	"gostartup/auth"
	"gostartup/handler"
	"gostartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connection := "root:@tcp(127.0.0.1:3306)/gostartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// memanggil db model repositor
	userRepository := user.NewRepository(db)

	// memanggil service
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// Handler atau controller logic
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)
	// api.GET("/users/fetch", userHandler.FetchUser)

	router.Run()
}
