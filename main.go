package main

import (
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

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()
	// user := user.User{
	// 	Name: "Test simpan",
	// }

	// userRepository.Save(user)

	// fmt.Println("Success Connect to database")

	// var users []user.User
	// length := len(users)

	// fmt.Println(length)

	// db.Find(&users)

	// length = len(users)
	// fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("=============================")
	// }

	// membuat router
	// router := gin.Default()
	// router.GET("/handler", handler)

	// jalankan router
	// router.Run()

}

// untuk membuat web, butuh func handler atau sebuah controller jika di laravel
// func handler(c *gin.Context) {
// 	connection := "root:@tcp(127.0.0.1:3306)/gostartup?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User

// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)

// }
