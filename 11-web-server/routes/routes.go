package routes

import (
	"github.com/Tyler-Liquornik/golang-fundamentals/11-web-server/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("events/:id", getEvent)
	server.GET("/events", getEvents)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)
	authenticated.POST("/events/:id/registrations", createRegistration)
	authenticated.DELETE("/events/:id/registrations", deleteRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
