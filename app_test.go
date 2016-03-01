package main

import (
	. "github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

const Cookie = "Set-Cookie"

func client(method, path, session string) *httptest.ResponseRecorder {
	gin.SetMode("test")
	app := NewApp()
	req, _ := http.NewRequest(method, path, nil)
	if len(session) != 0 {
		req.Header.Set(Cookie, session)
	}
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	return w
}

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("App api", func() {
		var session string

		g.It("Should return 200 on / ", func() {
			w := client("GET", "/", "")

			g.Assert(w.Code).Equal(200)
			session = w.HeaderMap.Get(Cookie)

		})

		g.It("Should return 200 on /slides.md ", func() {
			w := client("GET", "/slides.md", session)
			g.Assert(w.Code).Equal(200)
		})

		g.It("Should return 200 on PUT /slides.md ", func() {
			w := client("PUT", "/slides.md", session)
			g.Assert(w.Code).Equal(200)
		})

		g.It("Should works")

	})
}
