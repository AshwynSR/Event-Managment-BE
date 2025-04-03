package routes

import (
	"example/event-management/models"
	"example/event-management/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {

	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error while parsing data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error while Saving user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User was created!"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error while parsing data"})
		return
	}

	err = user.Validate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Login Successful!", "token": token})
}
