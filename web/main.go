package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	router.Static("/assets", "./web/assets")

	router.LoadHTMLGlob("./web/assets/html/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})

	router.POST("/create-gif", func(c *gin.Context) {
		width := c.PostForm("size-y")
		fmt.Println(width)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page..." + width,
		})
	})

	router.Run(":8080")
}
