package controllers

import (
	"github.com/gin-gonic/gin"
	"main.go/db"
	"main.go/middleware"
)

func Routes(r *gin.Engine) {

	r.POST("/signup", Signup)
	r.POST("/login", Login)

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.File("./static/login/signup.html")
	})

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.TokenAuthMiddleware())

	// Serve protected static files
	protected.GET("/home", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	protected.GET("/getAll.html", func(c *gin.Context) {
		c.File("./static/getAll.html")
	})
	protected.GET("/UpdateRecord.html", func(c *gin.Context) {
		c.File("./static/UpdateRecord.html")
	})
	protected.GET("/deleteRecord.html", func(c *gin.Context) {
		c.File("./static/deleteRecord.html")
	})
	protected.GET("/createTable.html", func(c *gin.Context) {
		c.File("./static/createTable.html")
	})

	r.POST("/stats", stats)
	r.GET("/display", DisplayAll)

	r.PUT("/update", db.UpdateRecord)
	r.DELETE("/delete", db.DeleteRecords)
	r.POST("/create-table-file", db.Createtable)
	r.POST("/create-table-users", db.CreateTableUser)
}
