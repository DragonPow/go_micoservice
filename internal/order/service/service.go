package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitService() {
	// Initialize your service here
	r := gin.Default()

	// Add your routes and middleware here

	// Connect to the database
	db, err := gorm.Open("mysql", "user:password@tcp(localhost:3306)/database_name?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	// Start the server
	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	InitService()
}
