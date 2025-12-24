package routes

import (
	"net/http"
	"rest-api-project/models"
	"rest-api-project/utils"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the requested data", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Sucessfully saved!", "User": user})

}

func login(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the requested data", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not login", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not login", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "token": token, "User": user})

}
