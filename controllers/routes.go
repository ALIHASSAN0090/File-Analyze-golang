package controllers

import (
	"github.com/gin-gonic/gin"
	"main.go/db"
)

func Routes(r *gin.Engine) {

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	r.GET("/index.html", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	r.GET("/test.html", func(c *gin.Context) {
		c.File("./static/test.html")
	})

	r.GET("/getAll.html", func(c *gin.Context) {
		c.File("./static/getAll.html")
	})

	r.GET("/UpdateRecord.html", func(c *gin.Context) {
		c.File("./static/UpdateRecord.html")
	})

	r.GET("/deleteRecord.html", func(c *gin.Context) {
		c.File("./static/deleteRecord.html")
	})

	r.GET("/createTable.html", func(c *gin.Context) {
		c.File("./static/createTable.html")
	})

	r.POST("/stats", stats)
	r.GET("/display", DisplayAll)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"working": "route is working"})
	})

	r.PUT("/update", db.UpdateRecord)
	r.DELETE("/delete", db.DeleteRecords)
	r.POST("/create-table", db.Createtable)
}
