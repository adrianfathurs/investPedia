package handler

import (
	"fmt"
	"investPedia/auth"
	"investPedia/helper"
	"investPedia/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Input is not valid", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account tfailed", "failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Problem on token JWT", "failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatterResponse(newUser, token)
	response := helper.APIResponse("Account has been registered", "success", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed, Input is not valid", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed, Input is not valid", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Problem on token JWT", "failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatterResponse(loggedinUser, token)
	response := helper.APIResponse("Login Success", "success", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
	// user memasukan email dan password
	// input ditangkap handler
	// mapping login struct terhadap input
	// di service akan mencari dalam bantuan repository user dengan email x
	// kalo ketemu mencocokan password
}

func (h *userHandler) CheckEmailAvailibility(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "something wrong in server"}
		response := helper.APIResponse("email already exist", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, "success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)

	// ada input dari user
	// inputan email akan di mapping kedalam struct
	// struct input di passing pada service
	// service laporan pada repository
	// repository akan melakukan pengecekan di dalam db
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed Upload File", "failed", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1
	filePath := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Request timeout, Failed Upload File", "failed", http.StatusRequestTimeout, data)
		c.JSON(http.StatusRequestTimeout, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, filePath)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Request timeout, Failed Upload File", "failed", http.StatusRequestTimeout, data)
		c.JSON(http.StatusRequestTimeout, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar is uploaded", "success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)

}
