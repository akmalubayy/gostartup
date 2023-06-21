package main

import (
	"gostartup/auth"
	"gostartup/campaign"
	"gostartup/handler"
	"gostartup/helper"
	"gostartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	campaignRepository := campaign.NewRepository(db)

	// memanggil service
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()

	// campaigns, _ := campaignService.GetCampaigns(0)

	// fmt.Println(len(campaigns))

	// Handler atau controller logic
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("/images", "./images/avatar")
	api := router.Group("api/v1")

	// POST
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	// GET
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)

	// api.GET("/users/fetch", userHandler.FetchUser)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				"error",
				nil,
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				"error",
				nil,
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				"error",
				nil,
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				"error",
				nil,
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}

// Langkah-langak kebutuhan middleware
// Ambil nilai header authorization : bearer tokennya
// dari header auth, ambil nilai tokennya saja
// validasi token menggunakan validateToken method
// tentukan tokennya valid atau tidak, jika valid lanjut, jika tidak unauthorized
// ambil user_id
// ambil user dari database berdasarkan user_id lewat service
// set context isinya data user
