package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	haikunator "github.com/atrox/haikunatorgo"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const sessionHeader = "slide-session"

func Header(c *gin.Context, key string) string {
	if values, _ := c.Request.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}

func BasicAuth() gin.HandlerFunc {
	realm := "Authorization Required"
	realm = "Basic realm=" + strconv.Quote(realm)
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	enabled := isEnabled(user, password)
	if enabled {
		log.Warn("Auth mode enabled")
		log.Warn(fmt.Sprintf("Visit http://%s:%s@0.0.0.0:8080", user, password))
	}
	return func(c *gin.Context) {
		header := Header(c, "Authorization")
		if enabled && header != authorizationHeader(user, password) {
			// Credentials doesn't match, we return 401 and abort handlers chain.
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}

func isEnabled(user, password string) bool {
	switch {
	case user == "":
		return false
	case password == "":
		return false
	default:
		return true
	}
}

func authorizationHeader(user, password string) string {
	base := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
}

func NewApp() *gin.Engine {

	r := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions(sessionHeader, store))
	r.Use(BasicAuth())

	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {

		fname := c.Param("name")
		log.WithFields(log.Fields{
			"name": fname,
		}).Info("Restore?")

		haikunator := haikunator.New()
		haikunator.TokenLength = 0
		name := haikunator.Haikunate()
		path := fmt.Sprintf("slides/%s.md", name)
		log.WithFields(log.Fields{
			"path": path,
		}).Info("A new session")
		session := sessions.Default(c)
		session.Set("name", path)
		session.Save()

		c.HTML(200, "index.tmpl", gin.H{
			"pubTo": path,
		})
	})

	mustHaveSession := func(c *gin.Context) (string, error) {
		session := sessions.Default(c)
		val := session.Get("name")
		emptySession := errors.New("Emtpy session")
		if val == nil {
			c.String(400, "No context")
			return "", emptySession
		}
		log.WithFields(log.Fields{
			"path": val,
		}).Info("Got session")
		path, ok := val.(string)
		if !ok {
			c.String(400, "No context")
			return "", emptySession
		}
		return path, nil
	}

	r.GET("/slides.md", func(c *gin.Context) {
		path, err := mustHaveSession(c)
		if err != nil {
			return
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
		path, err := mustHaveSession(c)
		if err != nil {
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

	r.GET("/stash", func(c *gin.Context) {
		files, err := ioutil.ReadDir("slides")
		if err != nil {
			log.Fatal(err)
		}
		var stash []string
		for _, file := range files {
			stash = append(stash, file.Name())
		}
		c.HTML(200, "stash.tmpl", gin.H{
			"stash": stash,
		})
	})

	r.GET("/stash/edit/:name", func(c *gin.Context) {

		name := c.Param("name")
		log.WithFields(log.Fields{
			"name": name,
		}).Info("Restore session?")

		if strings.HasSuffix(name, ".md") {
			name = name[0 : len(name)-3]
		}
		path := fmt.Sprintf("slides/%s.md", name)
		session := sessions.Default(c)
		session.Set("name", path)
		session.Save()

		c.HTML(200, "index.tmpl", gin.H{
			"pubTo": path,
		})
	})

	r.GET("/published/slides/:name", func(c *gin.Context) {

		name := c.Param("name")
		log.WithFields(log.Fields{
			"name": name,
		}).Info("Published")

		if strings.HasSuffix(name, ".md") {
			name = name[0 : len(name)-3]
		}
		path := fmt.Sprintf("slides/%s.md", name)
		session := sessions.Default(c)
		session.Set("name", path)
		session.Save()
		c.HTML(200, "slides.tmpl", gin.H{
			"pubTo": path,
		})
	})

	return r

}

func main() {
	r := NewApp()
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		envPort := os.Getenv("PORT")
		if len(envPort) > 0 {
			port = envPort
		}
	}
	log.Info("Started http://0.0.0.0:8080")
	r.Run(fmt.Sprintf(":%s", port))
}
