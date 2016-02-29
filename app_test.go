package main

import (
	. "github.com/franela/goblin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(method, path string) *httptest.ResponseRecorder {
  app := NewApp()
  req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	return w
}

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("App api", func() {

		g.It("Should return 200 on / ", func() {
      w := performRequest("GET", "/")
			g.Assert(w.Code).Equal(200)
		})

		g.It("Should return 200 on /slides.md ", func() {
      w := performRequest("GET", "/slides.md")
			g.Assert(w.Code).Equal(200)
		})

		g.It("Should return 200 on PUT /slides.md ", func() {
      w := performRequest("PUT", "/slides.md")
			g.Assert(w.Code).Equal(200)
		})

		g.It("Should works")

	})
}
