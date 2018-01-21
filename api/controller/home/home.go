package home

import (
	"net/http"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Get("/", Index)
}

// Index is the landing page. Will redirect to Home.Index if user is already logged
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	v := c.View.New("home/index")

	v.Render(w, r)
}
