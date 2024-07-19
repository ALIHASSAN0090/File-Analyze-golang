package controllers

import (
	"github.com/gin-gonic/gin"
	"main.go/db"
)

func Routes(r *gin.Engine) {
	r.GET("/", DisplayAll)
	r.POST("/", Stats)
	r.PUT("/", db.UpdateRecord)
	r.DELETE("/delete", db.DeleteRecords)
}
