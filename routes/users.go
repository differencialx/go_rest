package routes

import (
	"go_rest/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	err = user.Save()
	if err != nil {
		log.Fatal(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}