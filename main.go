package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var DB = make(map[string]string)

func NewApp() *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
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

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "value": value})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	return r

}

func main() {
	r := NewApp()
	r.Run(":8080")
}
