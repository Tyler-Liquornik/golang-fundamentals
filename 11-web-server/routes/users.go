package routes

import (
	"net/http"

	"github.com/Tyler-Liquornik/golang-fundamentals/11-web-server/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
