package public

import (
	"net/http"
	"yana/core/session"

	"github.com/Dlacreme/httpd/back/wesult"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/router"
)

// LoadRoutes the routes.
func LoadRoutes() {
	router.Get("/", Index)
	router.Post("/", Index)

	router.Post("/login", SignIn)
	router.Get("/logout", LogOut)
}

// Index is the landing page. Will redirect to Home.Index if user is already logged
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	if session.Me(&c).Error == nil {
		http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
	}

	v := c.View.New("public/index")
	res := wesult.New(nil, nil)
	res.RenderView(w, r, v)
}
