package main

import (
	"net/http"
	"strconv"

	"go-rest-api/db"
	"go-rest-api/models"
	"go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
	server := gin.Default()
	
	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
