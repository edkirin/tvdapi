package main

import (
	"fmt"

	"tvdapi/api"
	"tvdapi/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const SERVER_HOST = "localhost"
const SERVER_PORT = "5000"

func initRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()
	routes := router.Group("/")
	{
		routes.GET("/ping", api.HandlePing)
		routes.GET("/companies", api.HandleCompanyList(db))
	}
	return router
}

func main() {
	db := database.InitDB()
	router := initRouter(db)
	router.Run(fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT))
}
