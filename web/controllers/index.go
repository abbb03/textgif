package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/abbb03/textgif/web/models"
	"github.com/gin-gonic/gin"
)

func InternalError(c *gin.Context, err error) {
	c.HTML(http.StatusInternalServerError, "index.html", gin.H{
		"content": err.Error(),
	})
}

func CreateImage(c *gin.Context) {
	width := c.PostForm("size-x")
	height := c.PostForm("size-y")
	text := c.PostForm("input-text")
	gif, err := models.CreateGifFile(text, width, height)
	if err != nil {
		InternalError(c, err)
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

func GetFile(c *gin.Context) {
	name := c.Param("name")

	fileBytes, err := ioutil.ReadFile("./tmp/" + name)
	if err != nil {
		InternalError(c, err)
		return
	}
	c.Data(http.StatusOK, "image/gif", fileBytes)
}
