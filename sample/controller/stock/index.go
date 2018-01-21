package stock

import (
	"net/http"

	"yana/core/stock/report"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/router"
)

// LoadRoutes the routes.
func LoadRoutes() {
	router.Get("/stock", Index)
	router.Post("/stock/import", Import)

	router.Get("/stock/create", CreateTemplate)
}

// Index is the landing page. Will redirect to Home.Index if user is already logged
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("stock/index")

	res := report.Report(&c, "")

	res.RenderView(w, r, v)
}
