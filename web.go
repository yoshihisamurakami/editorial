package main

import (
	"editorial/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", controller.Index)
	router.GET("/articles/:id", controller.Show)
	router.Static("/css/", "./public/css/")
	router.Run()
}
