package main

import (
	"fmt"
	"github.com/atrox/haikunatorgo"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func NewApp() *gin.Engine {

	r := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("templates/index.tmpl")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		haikunator := haikunator.NewHaikunator()
		haikunator.TokenLength = 0
		name := haikunator.Haikunate()
		path := fmt.Sprintf("slides/%s.md", name)
		session := sessions.Default(c)
		session.Set("name", path)
		session.Save()

		c.HTML(200, "users/index.tmpl", gin.H{
			"pubTo": path,
		})
	})

	r.GET("/slides.md", func(c *gin.Context) {
		session := sessions.Default(c)
		val := session.Get("name")
		path, ok := val.(string)
		if !ok {
			c.String(400, "No context")
		}
		if _, err := os.Stat(path); err != nil {
			// coppy sapmle markdown file to the path
			body, err := ioutil.ReadFile("initial-slides.md")
			if err != nil {
				panic(err)
			}
			ioutil.WriteFile(path, body, 0644)
			c.String(200, string(body))
			return
		}

		body, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		c.String(200, string(body))
	})

	r.PUT("/slides.md", func(c *gin.Context) {
		session := sessions.Default(c)
		val := session.Get("name")
		path, ok := val.(string)
		if !ok {
			c.String(400, "No context")
		}
		body, _ := ioutil.ReadAll(c.Request.Body)
		ioutil.WriteFile(path, body, 0644)
		c.String(200, "")
	})

	return r

}

func main() {
	r := NewApp()
	r.Run(":8080")
}
