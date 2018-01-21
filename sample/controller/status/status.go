// Package status provides all the error pages like 404, 405, 500, 501,
// and the page when a CSRF token is invalid.
package status

import (
	"net/http"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/router"
	"github.com/Dlacreme/httpd/back/werror"
	"github.com/Dlacreme/httpd/back/wesult"
)

// Load the routes.
func Load() {
	router.MethodNotAllowed(Error405)
	router.NotFound(Error404)
}

// Error404 - Page Not Found.
func Error404(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("status/error")
	res := wesult.New(nil, werror.New(http.StatusNotFound, "Not Found."))
	res.RenderView(w, r, v)
}

// Error405 - Method Not Allowed.
func Error405(allowedMethods string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c := flight.Context(w, r)
		v := c.View.New("status/error")
		res := wesult.New(nil, werror.New(http.StatusMethodNotAllowed, "Not allowed"))
		res.RenderView(w, r, v)
	}
}

// Error500 - Internal Server Error.
func Error500(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("status/error")
	res := wesult.New(nil, werror.New(http.StatusInternalServerError, "Internal Server Error"))
	res.RenderView(w, r, v)
}

// Error501 - Not Implemented.
func Error501(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("status/error")
	res := wesult.New(nil, werror.New(http.StatusNotImplemented, "Not ready."))
	res.RenderView(w, r, v)
}

// InvalidToken shows a page in response to CSRF attacks.
func InvalidToken(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("status/error")
	res := wesult.New(nil, werror.New(http.StatusUnauthorized, "Invalid Token."))
	res.RenderView(w, r, v)
}
