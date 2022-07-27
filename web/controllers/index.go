package controllers

import (
	"net/http"

	"github.com/abbb03/textgif/web/models"
	"github.com/gin-gonic/gin"
)

func CreateImage(c *gin.Context) {
	width := c.PostForm("size-x")
	height := c.PostForm("size-y")
	text := c.PostForm("input-text")
	gif, err := models.GetGif(text, width, height)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"content": err.Error(),
		})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": "This is an index page...",
		"gif":     gif,
	})
}

func LoadIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"content": "This is an index page...",
	})
}
