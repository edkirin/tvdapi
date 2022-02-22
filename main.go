package main

import (
	// "fmt"

	"fmt"

	"tvdapi/api"
	"tvdapi/database"

	"github.com/gin-gonic/gin"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

const SERVER_HOST = "localhost"
const SERVER_PORT = "5000"

func initRouter() *gin.Engine {
	// router := gin.Default()
	router := gin.New()

	routes := router.Group("/")
	{
		routes.GET("/ping", api.HandlePing)
	}

	return router
}

func main() {
	// connectionString :=
	// 	"host=localhost " +
	// 		"user=televend " +
	// 		"password=televend " +
	// 		"dbname=televend " +
	// 		"port=5432 " +
	// 		"sslmode=disable " +
	// 		"TimeZone=Europe/Zagreb"
	// db, _ := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	// var companies []orm.Company

	// result := db.Limit(10).Order("ID").Find(&companies)

	// for _, company := range companies {
	// 	fmt.Printf(
	// 		"%d: %s (%s)\n",
	// 		company.ID,
	// 		company.Caption,
	// 		company.Email,
	// 	)
	// }

	// fmt.Printf("result.RowsAffected: %d\n", result.RowsAffected)
	// fmt.Printf("len(venues): %d\n", len(companies))

	db := database.InitDB()
	// var companies []orm.Company

	// db.Limit(10).Order("ID").Find(&companies)

	// for _, company := range companies {
	// 	fmt.Printf(
	// 		"%d: %s (%s)\n",
	// 		company.ID,
	// 		company.Caption,
	// 		company.Email,
	// 	)
	// }

	// router := gin.Default()
	router := gin.New()
	routes := router.Group("/")
	{
		routes.GET("/ping", api.HandlePing)
		routes.GET("/companies", api.HandleCompanyList(db))
	}

	router.Run(fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT))
}
