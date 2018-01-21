package social

import (
	"fmt"
	"net/http"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/blue-jay/core/router"
)

func LoadRoutes() {
	router.Get("/social", Index)
	router.Post("/social", Index)

	router.Get("/news", GetNews)
	router.Get("/articles", GetArticles)
	router.Get("/events", GetEvents)

	router.Get("/news/create", LoadCreateNews)
	router.Post("/news/create", CreateNews)
}

func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	fmt.Printf("Not implemented: %v\n", c)

	return
}
