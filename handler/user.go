package handler

import (
	"fmt"
	"gostartup/auth"
	"gostartup/helper"
	"gostartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{
		userService,
		authService,
	}
}

func (handler *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct di atas akan dipassing sebagai paramater service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		// gin.H adalah map
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"User register failed!",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	NewUser, err := handler.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse(
			"User register failed!",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := handler.authService.GenerateToken(NewUser.ID)

	if err != nil {
		response := helper.APIResponse(
			"Your token is invalid",
			http.StatusBadRequest,
			"error",
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(NewUser, token)

	response := helper.APIResponse(
		"Account has been registered",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) Login(c *gin.Context) {
	// Step-setep
	// user memasukan input (email dan password)
	// mapping dari input user ke input struct
	// iput struct passing ke service
	// di service akan mencari dg bantuan repository, mencari user dan email inputan
	// if find, pencocokan password
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"User Login Failed!",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := handler.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse(
			"User Login Failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)

		return
	}

	token, err := handler.authService.GenerateToken(loggedinUser.ID)

	if err != nil {
		response := helper.APIResponse(
			"User Login Failed",
			http.StatusBadRequest,
			"error",
			nil,
		)

		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helper.APIResponse(
		"User Login Success, Welcome"+" "+loggedinUser.Name,
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)

}

func (handler *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Oops, your email already taken!",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)

		c.JSON(http.StatusUnprocessableEntity, response)

		return
	}

	isEmailAvailable, err := handler.userService.IsEmailAvailable(input)

	if err != nil {

		errorMessage := gin.H{"errors": "Something Missing"}
		response := helper.APIResponse(
			"Oops, your email already taken!",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage,
		)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	var metaMessage string

	if isEmailAvailable {
		metaMessage = "Congratulation! your email is avaliable"
	} else {
		metaMessage = "Email is already taken!"
	}

	response := helper.APIResponse(
		metaMessage,
		http.StatusOK,
		"success",
		data,
	)

	c.JSON(http.StatusOK, response)

}

func (handler *userHandler) UploadAvatar(c *gin.Context) {
	// input dari user
	// simpan gambar dalam folder
	// di service panggil repository
	// JWT (user login id 1)
	// repo ambil data user yang id = 1
	// repo update data user dan simpan lokasi file

	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload avatar image",
			http.StatusBadRequest,
			"error",
			data,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// harusnya dapat dari JWT
	userID := 1

	// path := "images/avatar/" + file.Filename

	path := fmt.Sprintf("images/avatar/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload avatar image",
			http.StatusBadRequest,
			"error",
			data,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = handler.userService.SaveAvatar(userID, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload avatar image",
			http.StatusBadRequest,
			"error",
			data,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse(
		"Succes uploaded Avatar",
		http.StatusOK,
		"success",
		data,
	)

	c.JSON(http.StatusOK, response)

}

func (handler *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentYser").(user.User)
	formatter := user.FormatUser(currentUser, "")
	response := helper.APIResponse(
		"Successfuly fetch user data",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}
