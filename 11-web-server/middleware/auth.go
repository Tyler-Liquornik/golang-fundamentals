package middleware

import (
	"net/http"

	"github.com/Tyler-Liquornik/golang-fundamentals/11-web-server/utils"
	"github.com/gin-gonic/gin"
)

// Authenticate is middleware to authenticate requests (i.e., return 401 if needed).
func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		// AbortWithStatusJSON because as auth middleware we want the request to not reach the core server.
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	// Add the user ID to the context for later use.
	ctx.Set("userId", userId)

	// Cleanly hand off to the next request handler / middleware.
	ctx.Next()
}
