package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main.go/controllers"
	"main.go/db"
	_ "main.go/db"
	_ "main.go/docs"
)

// @title File Analyzer in golang
// @description A app that analyzes the contents of a text file.
// @host localhost:3000
// @BasePath /
func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	_, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	controllers.Routes(r)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
