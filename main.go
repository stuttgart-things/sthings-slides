package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var DB = make(map[string]string)

func NewApp() *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("templates/index.tmpl")
  r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "users/index.tmpl", gin.H{
			"pubTo": "Users",
		})
	})

	r.GET("/slides.md", func(c *gin.Context) {
		body, err := ioutil.ReadFile("initial-slides.md")
		if err != nil {
			panic(err)
		}
		c.String(200, string(body))
	})

	r.PUT("/slides.md", func(c *gin.Context) {
		c.String(403, "")
	})

	return r

}

func main() {
	r := NewApp()
	r.Run(":8080")
}
