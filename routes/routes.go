package routes

import (
	"go-rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){

	server.GET("/events", getEvents) 
	server.GET("/events/:id", getEvent)
	
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) // Runs given middleware before other events for the ones in the group
	authenticated.POST("/events",  createEvent) 
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// server.POST("/events", middlewares.Authenticate, createEvent) // Multiple request handlers executed from left to right

	server.POST("/signup", signup)
	server.POST("/login", login)
}