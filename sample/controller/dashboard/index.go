package dashboard

import (
	"net/http"
	"yana/core/middleware/acl"

	"github.com/Dlacreme/httpd/back/wesult"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/router"
)

// LoadRoutes the routes.
func LoadRoutes() {
	router.Get("/dashboard", Index, acl.ForMember)
}

// Index is the landing page. Will redirect to Home.Index if user is already logged
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("dashboard/index")

	res := wesult.New(nil, nil)

	res.RenderView(w, r, v)
}
