package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gin-gonic/gin"
)

const Cookie = "Set-Cookie"

func client(method, path, cookie string) *httptest.ResponseRecorder {
	gin.SetMode("test")
	app := NewApp()
	req, _ := http.NewRequest(method, path, nil)
	if len(cookie) != 0 {
		req.Header.Set("Cookie", cookie)
	}
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	return w
}

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("App api", func() {
		var cookie string

		g.It("Should return 200 on / ", func() {
			w := client("GET", "/", "")

			g.Assert(w.Code).Equal(200)
			cookie = w.HeaderMap.Get(Cookie)
		})

		g.It("Should return 200 on /slides.md ", func() {
			w := client("GET", "/slides.md", cookie)
			g.Assert(w.Code).Equal(200)
		})

		g.It("Should return 200 on PUT /slides.md ", func() {
			w := client("PUT", "/slides.md", cookie)
			g.Assert(w.Code).Equal(200)
		})

		g.It("Should return list of slides ", func() {
			w := client("GET", "/stash", cookie)
			g.Assert(w.Code).Equal(200)
		})

		g.It("Should return specific slide in preview", func() {
			w := client("GET", "/published/shy-cell.md", cookie)
			g.Assert(w.Code).Equal(200)
		})

		g.It("Should return specific slide in edit mode", func() {
			w := client("GET", "/stash/edit/shy-cell.md", cookie)
			g.Assert(w.Code).Equal(200)
		})

		g.It("Should works")

	})
}
