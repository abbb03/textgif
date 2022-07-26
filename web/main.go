package main

import (
	"html/template"
	"strings"

	"github.com/abbb03/textgif/web/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.Static("/assets", "./web/assets")

	router.LoadHTMLGlob("./web/assets/html/*.html")

	router.GET("/", controllers.LoadIndex)

	router.POST("/create-gif", controllers.CreateImage)

	router.Run(":8080")
}
