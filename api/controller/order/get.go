package order

import (
	"net/http"

	ct "yana/api/common/controller"
	"yana/core/order"

	"github.com/Dlacreme/httpd/back/flight"
)

func Get(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	search := r.URL.Query().Get("search")

	ct.Final(w, order.Get(&c, search))
}
