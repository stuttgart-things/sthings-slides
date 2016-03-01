package main

import (
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	haikunator "github.com/atrox/haikunatorgo"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const SessionHeader = "slide-session"

func NewApp() *gin.Engine {

	r := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions(SessionHeader, store))

	r.LoadHTMLGlob("templates/index.tmpl")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		haikunator := haikunator.NewHaikunator()
		haikunator.TokenLength = 0
		name := haikunator.Haikunate()
		path := fmt.Sprintf("slides/%s.md", name)
		log.WithFields(log.Fields{
			"path": path,
		}).Info("A new session")
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
		if val == nil {
			c.String(400, "No context")
			return
		}
		log.WithFields(log.Fields{
			"path": val,
		}).Info("Got session")
		path, ok := val.(string)
		if !ok {
			c.String(400, "No context")
		}
		if _, err := os.Stat(path); err != nil {
			// copy sample markdown file to the path
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
		if val == nil {
			c.String(400, "No context")
			return
		}
		log.WithFields(log.Fields{
			"path": val,
		}).Info("Got session")
		path, ok := val.(string)
		if !ok {
			c.String(400, "No context")
			return
		}
		body, _ := ioutil.ReadAll(c.Request.Body)
		ioutil.WriteFile(path, body, 0644)
		log.WithFields(log.Fields{
			"size": len(body),
			"file": path,
		}).Info("Wrote to file")
		c.String(200, "")
	})

	return r

}

func main() {
	r := NewApp()
	r.Run(":8080")
}
