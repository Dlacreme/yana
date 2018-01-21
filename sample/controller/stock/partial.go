package stock

import (
	"net/http"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/wesult"
)

// CreateTemplate return partial to create or update new stock
func CreateTemplate(w http.ResponseWriter, r *http.Request) {

	c := flight.Context(w, r)
	v := c.View.New("stock/update")
	res := wesult.New(nil, nil)

	res.Partial(w, r, v)
}
